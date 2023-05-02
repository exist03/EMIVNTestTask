package handlers

import (
	"EMIVNTestTask/internal/keyboards"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	"github.com/vitaliy-ukiru/fsm-telebot/storages/memory"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
)

const (
	adminCommands     = "- admin create_card [cardID] [bankInfo] {LimitInfo}\n- admin connect_card [cardID] [owner]\n- admin create_shogun [Nickname] [TG username]\n- admin create_daimyo [Nickname] [TG username]\n- admin create_samurai [Nickname] [TG username]\n- admin create_collector [Nickname] [TG username]\n- admin set_daimyo_owner [daimyo nickname] [shogun nickname]\n- admin set_samurai_owner [samurai nickname] [daimyo nickname]\n- admin get_shogun_info [shogunID]\n- admin get_daimyo_info [daimyoID]\n- admin get_samurai_info [samuraiID]\n- admin get_collector_info [collectorID]"
	shogunCommands    = "- shogun daimyos - Просмотр подчиненных(Даймё)\n- shogun samurais [daimyo nickname] - Просмотр подчиненных самураев у конкретного даймё\n- shogun create [cardID] [bankInfo] {LimitInfo}\n- shogun connect [cardID] [owner] - Привезка карты к даймё по нику"
	daimyoCommands    = "- daimyo samurais\n- daimyo set [cardID] [balance] - Остаток на карте\n- daimyo cards\n- daimyo application [cardID] [value] - Создание заявки на пополнение карты[cardID] до суммы [value]"
	samuraiCommands   = "- samurai turnover [value] - оборот за смену"
	collectorCommands = "- collector show - Показать запросы на пополнение\n- collector apply [cardID] [value] - Выполнить запрос на пополнение"
)

var (
	InputSG           = fsm.NewStateGroup("reg")
	BeginShogunState  = InputSG.New("action")
	InputAgeState     = InputSG.New("age")
	InputHobbyState   = InputSG.New("hobby")
	InputConfirmState = InputSG.New("confirm")
)

func InitHandlers(bot *tele.Bot, db *sql.DB, manager *fsm.Manager) {
	bot.Use(middleware.AutoRespond())
	b := bot.Group()
	storage := memory.NewStorage()
	defer storage.Close()
	manager := fsm.NewManager(b, storage)
	beginHandlers(b, manager)

	//b.Handle(tele.OnText, func(c tele.Context) error {
	//	sl := strings.Split(c.Text(), " ")
	//	senderID := c.Sender().Username
	//	switch sl[0] {
	//	case "admin":
	//		if !validAdmin(db, senderID) {
	//			return c.Send("У вас нет прав на эту комманду")
	//		}
	//		return c.Send(initAdminHandlers(sl[1:], db, c.Sender().Username))
	//	case "shogun":
	//		if !validShogun(db, senderID) {
	//			return c.Send("У вас нет прав на эту комманду")
	//		}
	//		return c.Send(initShogunHandlers(sl[1:], db, c.Sender().Username), keyboards.ShogunKB())
	//	case "daimyo":
	//		if !validDaimyo(db, senderID) {
	//			return c.Send("У вас нет прав на эту комманду")
	//		}
	//		return c.Send(initDaimyoHandlers(sl[1:], db, c.Sender().Username))
	//	case "samurai":
	//		if !validSamurai(db, senderID) {
	//			return c.Send("У вас нет прав на эту комманду")
	//		}
	//		return c.Send(initSamuraiHandlers(sl[1:], db, c.Sender().Username))
	//	case "collector":
	//		if !validCollector(db, senderID) {
	//			return c.Send("У вас нет прав на эту комманду")
	//		}
	//		return c.Send(initCollectorHandlers(sl[1:], db, c.Sender().Username))
	//	}
	//	return c.Send("Incorrect message")
	//})
}

func beginHandlers(b *tele.Group, manager *fsm.Manager) {
	b.Handle("/start", onStart())
	manager.Bind("/state", fsm.AnyState, func(c tele.Context, state fsm.FSMContext) error {
		s := state.State()
		return c.Send(s.String())
	})
	manager.Bind(&keyboards.BtnDaimyo, fsm.AnyState, onStartDaimyo())
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
	//b.Handle(&keyboards.BtnCollector, func(c tele.Context) error {
	//	return c.Send(collectorCommands)
	//})
}

func onStart() tele.HandlerFunc {
	return func(c tele.Context) error {
		log.Println("new user", c.Sender().ID)
		return c.Send("Choose", keyboards.StartKB())
	}
}

func onStartDaimyo() fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		state.Set(BeginShogunState)
		return c.Send("Выберите действие", keyboards.DaimyoKB())
	}
}

func validAdmin(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Admins WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
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
func validSamurai(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Samurais WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
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
