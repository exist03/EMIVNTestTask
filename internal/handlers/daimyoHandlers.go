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
	cardIDInputState = InputSG.New("cardIDInputState")
	AmountInputState = InputSG.New("AmountInputState")
)

func initDaiyoHandlers(manager *fsm.Manager, db *sql.DB) {
	manager.Bind(&keyboards.BtnDaimyo, fsm.DefaultState, onStartDaimyo())
	manager.Bind(&keyboards.BtnShowSamurais, BeginDaimyoState, showSamurais(db))
	manager.Bind(&keyboards.BtnShowCards, BeginDaimyoState, showCards(db))
	manager.Bind(&keyboards.BtnRemainingFunds, BeginDaimyoState, remainingFunds)
	manager.Bind(tele.OnText, cardIDInputState, onInputCard)
	manager.Bind(tele.OnText, AmountInputState, onInputRemainnings(db))
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

func onInputRemainnings(db *sql.DB) fsm.Handler {
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
