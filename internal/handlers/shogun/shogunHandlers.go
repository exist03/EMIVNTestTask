package shogun

import (
	"EMIVNTestTask/internal/handlers/states"
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/internal/users"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
)

var (
	BeginShogunState   = states.InputSG.New("beginShogunState")
	daimyoIDInputState = states.InputSG.New("daimyoIDInputState")
	//card create
	OnInputCardCreateState = states.InputSG.New("onInputCardCreateState")
	OnInputBankState       = states.InputSG.New("OnInputBankState")
	OnInputLimitState      = states.InputSG.New("OnInputLimitState")
	//card connect
	OnInputCardState      = states.InputSG.New("OnInputCardState")
	OnInputCardOwnerState = states.InputSG.New("OnInputCardOwnerState")
)

func InitShogunHandlers(manager *fsm.Manager, db *sql.DB) {
	//start buttons
	manager.Bind(&keyboards.BtnShogun, fsm.DefaultState, onStartShogun(db))
	//report
	manager.Bind(&keyboards.BtnGetReport, BeginShogunState, onReport)
	//report
	initReportHandlers(manager, db)

	manager.Bind(&keyboards.BtnShowDaimyos, BeginShogunState, showDaimyos(db))
	manager.Bind(&keyboards.BtnShowDaimyoSamurais, BeginShogunState, getDaimyoSamurais)
	manager.Bind(tele.OnText, daimyoIDInputState, showDaimyoSamurais(db))
	//cards
	//card create
	manager.Bind(&keyboards.BtnCreateCard, BeginShogunState, OnCreateCard)
	manager.Bind(tele.OnText, OnInputCardCreateState, OnInputCardCreate)
	manager.Bind(tele.OnText, OnInputBankState, OnInputBank)
	manager.Bind(tele.OnText, OnInputLimitState, OnInputLimit(db))
	//card connect
	manager.Bind(&keyboards.BtnConnectCard, BeginShogunState, OnConnectCard)
	manager.Bind(tele.OnText, OnInputCardState, OnInputCardConnect)
	manager.Bind(tele.OnText, OnInputCardOwnerState, OnInputCardOwner(db))
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

// report
func onReport(c tele.Context, state fsm.FSMContext) error {
	go state.Set(reportState)
	return c.Send("Выберите", keyboards.ShogunReportPeriodKB())
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
func OnConnectCard(c tele.Context, state fsm.FSMContext) error {
	go state.Set(OnInputCardState)
	return c.Send("Введите карту которую хотите привязать")
}
func OnInputCardConnect(c tele.Context, state fsm.FSMContext) error {
	go state.Update("cardID", c.Text())
	go state.Set(OnInputCardOwnerState)
	return c.Send("Введите дайме к которому хотите привязать карту")
}
func OnInputCardOwner(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(fsm.DefaultState)
		cardModel := mysql.CardModel{DB: db}
		cardID, err := state.Get("cardID")
		if err != nil {
			log.Println(err)
			return c.Send("Произошла ошибка")
		}
		cardModel.SetOwner(cardID, c.Text())
		return c.Send("Карта привязана", keyboards.StartKB())
	}
}

// create card
func OnCreateCard(c tele.Context, state fsm.FSMContext) error {
	go state.Set(OnInputCardCreateState)
	return c.Send("Введите карту которую хотите создать")
}
func OnInputCardCreate(c tele.Context, state fsm.FSMContext) error {
	go state.Update("cardID", c.Text())
	go state.Set(OnInputBankState)
	return c.Send("Введите банк")
}
func OnInputBank(c tele.Context, state fsm.FSMContext) error {
	go state.Update("bank", c.Text())
	go state.Set(OnInputLimitState)
	return c.Send("Введите лимит")
}
func OnInputLimit(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(fsm.DefaultState)
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
		return c.Send("Карта создана", keyboards.StartKB())
	}
}
