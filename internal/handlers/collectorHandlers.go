package handlers

import (
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

var (
	BeginCollectorState = InputSG.New("startCollector")
	onInputCardIDState  = InputSG.New("onInputCardIDState")
	OnInputAmountState  = InputSG.New("OnInputAmountState")
)

//
//
//		return res

func initCollectorHandlers(manager *fsm.Manager, db *sql.DB) {
	manager.Bind(&keyboards.BtnCollector, fsm.DefaultState, onStartCollector(db))
	//show
	manager.Bind(&keyboards.BtnShowApplications, BeginCollectorState, onShow(db))
	//apply
	manager.Bind(&keyboards.BtnApplyApplication, BeginCollectorState, onApplyApplication)
	manager.Bind(tele.OnText, onInputCardIDState, onInputCardID)
	manager.Bind(tele.OnText, OnInputAmountState, onInputSum(db))
}

func onShow(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		CollectorModel := mysql.CollectorModel{DB: db}
		res := CollectorModel.ShowApplications()
		return c.Send(res, keyboards.CollectorKB())
	}
}

func onApplyApplication(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onInputCardIDState)
	return c.Send("Введите карту для пополнения")
}
func onInputCardID(c tele.Context, state fsm.FSMContext) error {
	go state.Update("cardID", c.Text())
	go state.Set(OnInputAmountState)
	return c.Send("Введите сумму")
}
func onInputSum(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		amount, err := strconv.ParseFloat(c.Text(), 64)
		if err != nil {
			return c.Send("Некорректный ввод. Повторите попытку")
		}
		cardID, err := state.Get("cardID")
		if err != nil {
			return c.Send("Произошла ошибка. Повторите попытку")
		}
		CollectorModel := mysql.CollectorModel{DB: db}
		res := CollectorModel.ApplyApplication(cardID, amount)
		state.Set(BeginCollectorState)
		return c.Send(res, keyboards.CollectorKB())
	}
}

func onStartCollector(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		if !validCollector(db, c.Sender().Username) {
			return c.Send("У вас нет прав")
		}
		state.Set(BeginCollectorState)
		return c.Send("Выберите действие", keyboards.CollectorKB())
	}
}
func validCollector(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Collectors WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
