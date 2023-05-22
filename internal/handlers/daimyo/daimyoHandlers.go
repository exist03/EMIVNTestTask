package daimyo

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
	BeginDaimyoState    = states.InputSG.New("startDaimyo")
	cardIDInputState    = states.InputSG.New("cardIDInputState")
	cardIDInputStateApp = states.InputSG.New("cardIDInputStateApp")
	AmountInputState    = states.InputSG.New("AmountInputState")
	AmountInputStateApp = states.InputSG.New("AmountInputStateApp")
)

func InitDaiyoHandlers(manager *fsm.Manager, db *sql.DB) {
	manager.Bind(&keyboards.BtnDaimyo, fsm.DefaultState, onStartDaimyo(db))

	//start buttons
	manager.Bind(&keyboards.BtnHierarchy, BeginDaimyoState, onHierarchy)
	manager.Bind(&keyboards.BtnReport, BeginDaimyoState, onReport)
	manager.Bind(&keyboards.BtnCardLimit, BeginDaimyoState, cardLimit(db))
	initHierarchyHandlers(manager, db)
	initReportHandlers(manager, db)
	//application
	manager.Bind(&keyboards.BtnCreateApplication, BeginDaimyoState, beginApp)
	manager.Bind(tele.OnText, cardIDInputStateApp, onInputCardApp)
	manager.Bind(tele.OnText, AmountInputStateApp, onInputAmountApp(db))
	//card limit
	manager.Bind(&keyboards.BtnHierarchy, BeginDaimyoState, onHierarchy)
	//remaining
	manager.Bind(&keyboards.BtnRemainingFunds, BeginDaimyoState, remainingFunds)
	manager.Bind(tele.OnText, cardIDInputState, onInputCard)
	manager.Bind(tele.OnText, AmountInputState, onInputAmount(db))
}
func onStartDaimyo(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		if !validDaimyo(db, c.Sender().Username) {
			return c.Send("У вас нет прав")
		}
		state.Set(BeginDaimyoState)
		return c.Send("Выберите действие", keyboards.DaimyoKB())
	}
}

func onHierarchy(c tele.Context, state fsm.FSMContext) error {
	go state.Set(hierarchyState)
	return c.Send("Выберите", keyboards.DaimyoHierarchyKB())
}
func onReport(c tele.Context, state fsm.FSMContext) error {
	go state.Set(reportState)
	return c.Send("Выберите", keyboards.DaimyoReportKB())
}
func cardLimit(db *sql.DB) fsm.Handler {
	cardModel := mysql.CardModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		res, _ := cardModel.CardLimit(c.Sender().Username)
		for _, v := range res {
			c.Send(v)
		}
		return nil
	}
}
func validTime() bool {
	currentHour := time.Now().Hour()
	return currentHour >= 8 && currentHour < 12
}

//	func showCards(db *sql.DB) fsm.Handler {
//		cardModel := mysql.CardModel{DB: db}
//		return func(c tele.Context, state fsm.FSMContext) error {
//			res, _ := cardModel.GetList(c.Sender().Username)
//			for _, v := range res {
//				c.Send(v)
//			}
//			return nil
//		}
//	}
func remainingFunds(c tele.Context, state fsm.FSMContext) error {
	go state.Set(cardIDInputState)
	return c.Send("Введите номер карты")
}
func onInputCard(c tele.Context, state fsm.FSMContext) error {
	cardID := c.Message().Text
	go state.Update("cardID", cardID)
	go state.Set(AmountInputState)
	return c.Send("Введите сумму")
}
func onInputAmount(db *sql.DB) fsm.Handler {
	cardModel := mysql.CardModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(BeginDaimyoState)
		cardID, err := state.Get("cardID")
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		amount, err := strconv.ParseFloat(c.Text(), 64)
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		res := cardModel.Update(cardID, amount)
		return c.Send(res, keyboards.DaimyoKB())
	}
}
func beginApp(c tele.Context, state fsm.FSMContext) error {
	go state.Set(cardIDInputStateApp)
	return c.Send("Введите номер карты")
}
func onInputCardApp(c tele.Context, state fsm.FSMContext) error {
	cardID := c.Message().Text
	go state.Update("cardID", cardID)
	go state.Set(AmountInputStateApp)
	return c.Send("Введите сумму")
}
func onInputAmountApp(db *sql.DB) fsm.Handler {
	daimyoModel := mysql.DaimyoModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(BeginDaimyoState)
		cardID, err := state.Get("cardID")
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		amount, err := strconv.ParseFloat(c.Text(), 64)
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		res := daimyoModel.InsertApp(c.Sender().Username, cardID, amount)
		return c.Send(res, keyboards.DaimyoKB())
	}
}
func validDaimyo(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Daimyo WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
