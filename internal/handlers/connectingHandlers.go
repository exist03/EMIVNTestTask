package handlers

import (
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
)

var (
	onConnectDaimyoUsernameState  = InputSG.New("onConnectDaimyoUsernameState")
	onConnectDaimyoOwnerState     = InputSG.New("onConnectDaimyoOwnerState")
	onConnectSamuraiUsernameState = InputSG.New("onConnectSamuraiUsernameState")
	onConnectSamuraiOwnerState    = InputSG.New("onConnectSamuraiOwnerState")
)

func initConnectingHandlers(manager *fsm.Manager, db *sql.DB) {
	//card
	manager.Bind(&keyboards.BtnConnectCard, onConnectState, onConnectCard)
	manager.Bind(tele.OnText, onInputCardState, onInputCardConnect)
	manager.Bind(tele.OnText, onInputCardOwnerState, onInputCardOwner(db))
	//daimyo
	manager.Bind(&keyboards.BtnDaimyo, onConnectState, onConnectDaimyo)
	manager.Bind(tele.OnText, onConnectDaimyoUsernameState, onConnectDaimyoUsername)
	manager.Bind(tele.OnText, onConnectDaimyoOwnerState, onConnectDaimyoOwner(db))
	//samurai
	manager.Bind(&keyboards.BtnSamurai, onConnectState, onConnectSamurai)
	manager.Bind(tele.OnText, onConnectSamuraiUsernameState, onConnectSamuraiUsername)
	manager.Bind(tele.OnText, onConnectSamuraiOwnerState, onConnectSamuraiOwner(db))
}

func onConnecting(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onConnectState)
	return c.Send("Выберите", keyboards.AdminConnectKB())
}

// daimyo
func onConnectDaimyo(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onConnectDaimyoUsernameState)
	return c.Send("Введите username дайме которого нужно привязать")
}
func onConnectDaimyoUsername(c tele.Context, state fsm.FSMContext) error {
	go state.Update("connectedDaimyo", c.Text())
	go state.Set(onConnectDaimyoOwnerState)
	return c.Send("Введите username сегуна к которому нужно привязать")
}
func onConnectDaimyoOwner(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(beginAdminState)
		daimyo, err := state.Get("connectedDaimyo")
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		daimyoModel := mysql.DaimyoModel{DB: db}
		daimyoModel.SetOwner(daimyo, c.Text())
		if err != nil {
			return c.Send("Произошла ошбика", keyboards.AdminKB())
		}
		return c.Send("Дайме привязан", keyboards.AdminKB())
	}
}

// samurai
func onConnectSamurai(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onConnectSamuraiUsernameState)
	return c.Send("Введите username самурая которого нужно привязать")
}
func onConnectSamuraiUsername(c tele.Context, state fsm.FSMContext) error {
	go state.Update("connectedSamurai", c.Text())
	go state.Set(onConnectSamuraiOwnerState)
	return c.Send("Введите username дайме к которому нужно привязать")
}
func onConnectSamuraiOwner(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(beginAdminState)
		samurai, err := state.Get("connectedSamurai")
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		samuraiModel := mysql.SamuraiModel{DB: db}
		samuraiModel.SetOwner(samurai, c.Text())
		if err != nil {
			return c.Send("Произошла ошбика", keyboards.AdminKB())
		}
		return c.Send("Дайме привязан", keyboards.AdminKB())
	}
}
