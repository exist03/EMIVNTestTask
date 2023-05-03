package handlers

import (
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/internal/users"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
)

var (
	daimyoIDInputState = InputSG.New("daimyoIDInputState")
	//card create
	onInputCardCreateState = InputSG.New("onInputCardCreateState")
	onInputBankState       = InputSG.New("onInputBankState")
	onInputLimitState      = InputSG.New("onInputLimitState")
	//card connect
	onInputCardState      = InputSG.New("onInputCardState")
	onInputCardOwnerState = InputSG.New("onInputCardOwnerState")
)

func initShogunHandlers(manager *fsm.Manager, db *sql.DB) {
	//start buttons
	manager.Bind(&keyboards.BtnShogun, fsm.DefaultState, onStartShogun(db))
	manager.Bind(&keyboards.BtnShowDaimyos, BeginShogunState, showDaimyos(db))
	manager.Bind(&keyboards.BtnShowDaimyoSamurais, BeginShogunState, getDaimyoSamurais)
	manager.Bind(tele.OnText, daimyoIDInputState, showDaimyoSamurais(db))
	//cards
	//card create
	manager.Bind(&keyboards.BtnCreateCard, BeginShogunState, onCreateCard)
	manager.Bind(tele.OnText, onInputCardCreateState, onInputCardCreate)
	manager.Bind(tele.OnText, onInputBankState, onInputBank)
	manager.Bind(tele.OnText, onInputLimitState, onInputLimit(db))
	//card connect
	manager.Bind(&keyboards.BtnConnectCard, BeginShogunState, onConnectCard)
	manager.Bind(tele.OnText, onInputCardState, onInputCardConnect)
	manager.Bind(tele.OnText, onInputCardOwnerState, onInputCardOwner(db))
}
func onStartShogun(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		if !validShogun(db, c.Sender().Username) {
			return c.Send("У вас нет прав")
		}
		state.Set(BeginShogunState)
		return c.Send("Выберите действие", keyboards.ShogunKB())
	}
}
func validShogun(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Shogun WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func showDaimyos(db *sql.DB) fsm.Handler {
	daimyoModel := mysql.DaimyoModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		res, _ := daimyoModel.GetList(c.Sender().Username)
		return c.Send(res)
	}
}

// samurais
func getDaimyoSamurais(c tele.Context, state fsm.FSMContext) error {
	go state.Set(daimyoIDInputState)
	return c.Send("Введите username дайме")
}
func showDaimyoSamurais(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(BeginShogunState)
		samuraiModel := mysql.SamuraiModel{DB: db}
		daimyoID := c.Message().Text
		res, _ := samuraiModel.GetList(daimyoID)
		return c.Send(res, keyboards.ShogunKB())
	}
}

// connect card
func onConnectCard(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onInputCardState)
	return c.Send("Введите карту которую хотите привязать")
}
func onInputCardConnect(c tele.Context, state fsm.FSMContext) error {
	go state.Update("cardID", c.Text())
	go state.Set(onInputCardOwnerState)
	return c.Send("Введите дайме к которому хотите привязать карту")
}
func onInputCardOwner(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(BeginShogunState)
		cardModel := mysql.CardModel{DB: db}
		cardID, err := state.Get("cardID")
		if err != nil {
			log.Println(err)
			return c.Send("Произошла ошибка")
		}
		cardModel.SetOwner(cardID, c.Text())
		return c.Send("Карта привязана", keyboards.ShogunKB())
	}
}

// create card
func onCreateCard(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onInputCardCreateState)
	return c.Send("Введите карту которую хотите создать")
}
func onInputCardCreate(c tele.Context, state fsm.FSMContext) error {
	go state.Update("cardID", c.Text())
	go state.Set(onInputBankState)
	return c.Send("Введите банк")
}
func onInputBank(c tele.Context, state fsm.FSMContext) error {
	go state.Update("bank", c.Text())
	go state.Set(onInputLimitState)
	return c.Send("Введите лимит")
}
func onInputLimit(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(BeginShogunState)
		cardModel := mysql.CardModel{DB: db}
		cardID, err := state.Get("cardID")
		if err != nil {
			log.Println(err)
			return c.Send("Произошла ошибка")
		}
		bank, err := state.Get("bank")
		if err != nil {
			log.Println(err)
			return c.Send("Произошла ошибка")
		}
		card := users.Card{
			ID:        cardID,
			Owner:     "1",
			BankInfo:  bank,
			LimitInfo: c.Text(),
			Balance:   c.Text(),
		}
		cardModel.Insert(card)
		return c.Send("Карта создана", keyboards.ShogunKB())
	}
}
