package samurai

import (
	"EMIVNTestTask/internal/handlers/states"
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
	"time"
)

var (
	BeginSamuraiState    = states.InputSG.New("startSamurai")
	OnInputTurnoverState = states.InputSG.New("OnInputTurnoverState")
	OnInputBankNameState = states.InputSG.New("OnInputBankNameState")
)

func InitSamuraiHandlers(manager *fsm.Manager, db *sql.DB) {
	manager.Bind(&keyboards.BtnSamurai, fsm.DefaultState, onStartSamurai(db))
	manager.Bind(&keyboards.BtnTurnover, BeginSamuraiState, onInputTurnover)
	manager.Bind(tele.OnText, OnInputBankNameState, onInputBankName)
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
	stmt := `SELECT COUNT(*) FROM Samurai WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func validTime() bool {
	currentHour := time.Now().Hour()
	return currentHour >= 8 && currentHour < 12
}

func onInputBankName(c tele.Context, state fsm.FSMContext) error {
	go state.Update("bankName", c.Text())
	go state.Set(OnInputTurnoverState)
	return c.Send("Введите сумму")
}
func onInputTurnover(c tele.Context, state fsm.FSMContext) error {
	if !validTime() {
		return c.Send("Внесение данных возможно с 8:00 до 12:00")
	}
	go state.Set(OnInputBankNameState)
	return c.Send("Введите банк")
}

func onTurnover(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		samuraiModel := mysql.SamuraiModel{DB: db}
		value, err := strconv.ParseFloat(c.Text(), 64)
		if err != nil {
			log.Println(err)
			return c.Send("Некорректное значение. Попробуйте еще раз", keyboards.SamuraiKB())
		}
		bank, err := state.Get("bankName")
		if err != nil {
			log.Println(err)
			return c.Send("Некорректное значение. Попробуйте еще раз", keyboards.SamuraiKB())
		}
		err = samuraiModel.SetTurnoverEnd(c.Sender().Username, bank, value)
		if err != nil {
			log.Println(err)
			go state.Set(BeginSamuraiState)
			return c.Send("Произошла ошибка", keyboards.SamuraiKB())
		}
		err = samuraiModel.SetTurnoverBegin(c.Sender().Username, bank, value)
		if err != nil {
			log.Println(err)
			go state.Set(BeginSamuraiState)
			return c.Send("Произошла ошибка", keyboards.SamuraiKB())
		}
		go state.Set(BeginSamuraiState)
		return c.Send("Данные записаны", keyboards.SamuraiKB())
	}
}
