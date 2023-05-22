package entryPoint

import (
	"EMIVNTestTask/internal/handlers/admin"
	"EMIVNTestTask/internal/handlers/collector"
	"EMIVNTestTask/internal/handlers/daimyo"
	"EMIVNTestTask/internal/handlers/samurai"
	"EMIVNTestTask/internal/handlers/shogun"
	"EMIVNTestTask/internal/keyboards"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
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
	admin.InitAdminHandlers(manager, db)
	shogun.InitShogunHandlers(manager, db)
	daimyo.InitDaiyoHandlers(manager, db)
	samurai.InitSamuraiHandlers(manager, db)
	collector.InitCollectorHandlers(manager, db)
}
func onStart() tele.HandlerFunc {
	return func(c tele.Context) error {
		log.Println("new user", c.Sender().ID)
		return c.Send("Выберите", keyboards.StartKB())
	}
}
