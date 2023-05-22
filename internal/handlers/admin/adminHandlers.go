package admin

import (
	"EMIVNTestTask/internal/handlers/states"
	"EMIVNTestTask/internal/keyboards"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
)

var (
	beginAdminState   = states.InputSG.New("startAdmin")
	onCreateState     = states.InputSG.New("onCreateState")
	onConnectState    = states.InputSG.New("onConnectState")
	onAdditionalState = states.InputSG.New("onAdditionalState")
)

func InitAdminHandlers(manager *fsm.Manager, db *sql.DB) {
	//start buttons
	manager.Bind(&keyboards.BtnAdmin, fsm.DefaultState, onStartAdmin(db))
	//create
	manager.Bind(&keyboards.BtnCreate, beginAdminState, onCreating)
	initCreatingHandlers(manager, db)
	//connect
	manager.Bind(&keyboards.BtnConnect, beginAdminState, onConnecting)
	initConnectingHandlers(manager, db)
	//additional
	manager.Bind(&keyboards.BtnAdditional, beginAdminState, onAdditional)
	initAdditionalHandlers(manager, db)
}
func onStartAdmin(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		if !validAdmin(db, c.Sender().Username) {
			return c.Send("У вас нет прав")
		}
		state.Set(beginAdminState)
		return c.Send("Выберите действие", keyboards.AdminKB())
	}
}
func validAdmin(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Admin WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
