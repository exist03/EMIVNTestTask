package handlers

import (
	"EMIVNTestTask/internal/keyboards"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
)

const (
	adminCommands     = "- admin create_card [cardID] [bankInfo] {LimitInfo}\n- admin connect_card [cardID] [owner]\n- admin create_shogun [Nickname] [TG username]\n- admin create_daimyo [Nickname] [TG username]\n- admin create_samurai [Nickname] [TG username]\n- admin create_collector [Nickname] [TG username]\n- admin set_daimyo_owner [daimyo nickname] [shogun nickname]\n- admin set_samurai_owner [samurai nickname] [daimyo nickname]\n- admin get_shogun_info [shogunID]\n- admin get_daimyo_info [daimyoID]\n- admin get_samurai_info [samuraiID]\n- admin get_collector_info [collectorID]"
	shogunCommands    = "- shogun daimyos - Просмотр подчиненных(Даймё)\n- shogun samurais [daimyo nickname] - Просмотр подчиненных самураев у конкретного даймё\n- shogun create [cardID] [bankInfo] {LimitInfo}\n- shogun connect [cardID] [owner] - Привезка карты к даймё по нику"
	samuraiCommands   = "- samurai turnover [value] - оборот за смену"
	collectorCommands = "- collector show - Показать запросы на пополнение\n- collector apply [cardID] [value] - Выполнить запрос на пополнение"
)

var (
	InputSG             = fsm.NewStateGroup("start")
	BeginAdminState     = InputSG.New("startAdmin")
	BeginShogunState    = InputSG.New("startShogun")
	BeginDaimyoState    = InputSG.New("startDaimyo")
	BeginSamuraiState   = InputSG.New("startSamurai")
	BeginCollectorState = InputSG.New("startCollector")
)

func InitHandlers(bot *tele.Group, db *sql.DB, manager *fsm.Manager) {
	bot.Handle("/start", onStart())
	manager.Bind("/state", fsm.AnyState, func(c tele.Context, state fsm.FSMContext) error {
		s := state.State()
		return c.Send(s.String())
	})
	manager.Bind(&keyboards.BtnCancel, fsm.AnyState, func(c tele.Context, state fsm.FSMContext) error {
		state.Set(fsm.DefaultState)
		return c.Send("Выберите", keyboards.StartKB())
	})
	initShogunHandlers(manager, db)
	initDaiyoHandlers(manager, db)

	//b.Handle(&keyboards.BtnAdmin, func(c tele.Context) error {
	//	return c.Send(adminCommands)
	//})
	//b.Handle(&keyboards.BtnShogun, func(c tele.Context) error {
	//	return c.Send(shogunCommands, keyboards.ShogunKB())
	//})
	//b.Handle(&keyboards.BtnDaimyo, func(c tele.Context) error {
	//	return c.Send(daimyoCommands, keyboards.DaimyoKB())
	//})
	//b.Handle(&keyboards.BtnSamurai, func(c tele.Context) error {
	//	return c.Send(samuraiCommands)
	//})

}
