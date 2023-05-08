package handlers

import (
	"EMIVNTestTask/internal/keyboards"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
)

var (
	InputSG = fsm.NewStateGroup("start")
)

func InitHandlers(bot *tele.Group, db *sql.DB, manager *fsm.Manager) {
	bot.Handle("/start", onStart())
	manager.Bind("/state", fsm.AnyState, func(c tele.Context, state fsm.FSMContext) error {
		s := state.State()
		return c.Send(s.String())
	})
	manager.Bind("/cancel", fsm.AnyState, func(c tele.Context, state fsm.FSMContext) error {
		state.Set(fsm.DefaultState)
		return c.Send("Состояние обнулено", keyboards.StartKB())
	})
	manager.Bind(&keyboards.BtnCancel, fsm.AnyState, func(c tele.Context, state fsm.FSMContext) error {
		state.Set(fsm.DefaultState)
		return c.Send("Выберите", keyboards.StartKB())
	})
	initAdminHandlers(manager, db)
	initShogunHandlers(manager, db)
	initDaiyoHandlers(manager, db)
	initSamuraiHandlers(manager, db)
	initCollectorHandlers(manager, db)
}
