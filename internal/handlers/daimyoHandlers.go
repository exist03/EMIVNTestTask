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
	BeginDaimyoState    = InputSG.New("startDaimyo")
	cardIDInputState    = InputSG.New("cardIDInputState")
	cardIDInputStateApp = InputSG.New("cardIDInputStateApp")
	AmountInputState    = InputSG.New("AmountInputState")
	AmountInputStateApp = InputSG.New("AmountInputStateApp")
)

func initDaiyoHandlers(manager *fsm.Manager, db *sql.DB) {
	//start buttons
	manager.Bind(&keyboards.BtnDaimyo, fsm.DefaultState, onStartDaimyo(db))
	manager.Bind(&keyboards.BtnShowSamurais, BeginDaimyoState, showSamurais(db))
	manager.Bind(&keyboards.BtnShowCards, BeginDaimyoState, showCards(db))
	//application
	manager.Bind(&keyboards.BtnCreateApplication, BeginDaimyoState, beginApp)
	manager.Bind(tele.OnText, cardIDInputStateApp, onInputCardApp)
	manager.Bind(tele.OnText, AmountInputStateApp, onInputAmountApp(db))
	//remaining
	manager.Bind(&keyboards.BtnRemainingFunds, BeginDaimyoState, remainingFunds)
	manager.Bind(tele.OnText, cardIDInputState, onInputCard)
	manager.Bind(tele.OnText, AmountInputState, onInputAmount(db))
}
func onStart() tele.HandlerFunc {
	return func(c tele.Context) error {
		log.Println("new user", c.Sender().ID)
		return c.Send("Выберите", keyboards.StartKB())
	}
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
func showSamurais(db *sql.DB) fsm.Handler {
	samuraiModel := mysql.SamuraiModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		res, _ := samuraiModel.GetList(c.Sender().Username)
		return c.Send(res)
	}
}
func showCards(db *sql.DB) fsm.Handler {
	cardModel := mysql.CardModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		res, _ := cardModel.GetList(c.Sender().Username)
		return c.Send(res)
	}
}
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
