package handlers

import (
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
)

var (
	BeginSamuraiState    = InputSG.New("startSamurai")
	OnInputTurnoverState = InputSG.New("OnInputTurnoverState")
)

func initSamuraiHandlers(manager *fsm.Manager, db *sql.DB) {
	manager.Bind(&keyboards.BtnSamurai, fsm.DefaultState, onStartSamurai(db))
	manager.Bind(&keyboards.BtnTurnover, BeginSamuraiState, onInputTurnover)
	manager.Bind(tele.OnText, OnInputTurnoverState, onTurnover(db))
}

func onStartSamurai(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		if !validSamurai(db, c.Sender().Username) {
			return c.Send("У вас нет прав")
		}
		state.Set(BeginSamuraiState)
		return c.Send("Выберите действие", keyboards.SamuraiKB())
	}
}

func validSamurai(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Samurais WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}

func onInputTurnover(c tele.Context, state fsm.FSMContext) error {
	go state.Set(OnInputTurnoverState)
	return c.Send("Введите сумму")
}

func onTurnover(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		samuraiModel := mysql.SamuraiModel{DB: db}
		value, err := strconv.ParseFloat(c.Text(), 64)
		if err != nil {
			log.Println(err)
			return c.Send("Некорректное значение. Введите еще раз", keyboards.SamuraiKB())
		}
		err = samuraiModel.SetTurnover(c.Sender().Username, value)
		if err != nil {
			log.Println(err)
			go state.Set(BeginSamuraiState)
			return c.Send("Произошла ошибка", keyboards.SamuraiKB())
		}
		go state.Set(BeginSamuraiState)
		return c.Send("Данные записаны", keyboards.SamuraiKB())
	}
}
