package handlers

import (
	"EMIVNTestTask/internal/keyboards"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
)

//	case "get_shogun_info": //admin get_shogun_info [shogunID]
//		if len(command) != 2 {
//			return "Wrong message"
//		}
//		shogunID := command[1]
//		shogunModel := mysql.ShogunModel{DB: db}
//		return shogunModel.Get(shogunID)
//	case "get_daimyo_info": //admin get_daimyo_info [daimyoID]
//		if len(command) != 2 {
//			return "Wrong message"
//		}
//		daimyoID := command[1]
//		daimyoModel := mysql.DaimyoModel{DB: db}
//		return daimyoModel.Get(daimyoID)
//	case "get_samurai_info": //admin get_samurai_info [samuraiID]
//		if len(command) != 2 {
//			return "Wrong message"
//		}
//		samuraiID := command[1]
//		samuraiModel := mysql.SamuraiModel{DB: db}
//		return samuraiModel.Get(samuraiID)
//	case "get_collector_info": //admin get_collector_info [collectorID]
//		if len(command) != 2 {
//			return "Wrong message"
//		}
//		collectorID := command[1]
//		collectorModel := mysql.CollectorModel{DB: db}
//		return collectorModel.Get(collectorID)
//	case "report_samurai": //admin report_samurai [dd.mm.yy] [samuraiID]
//		report := mysql.ReportModel{DB: db}
//		samuraiID := command[2]
//		time, _ := time2.Parse("2006-01-02", command[1])
//		return report.Samurais(samuraiID, time)
//		//баланс на начало смены(8:00 dd.mm.yyyy)
//		//сумма поступлений в течении смены
//		//сумма списаний в течении смены
//		//баланс на конец смены (8:00 dd+1.mm.yyyy)
//	}
//	return "Something went wrong"
//}

var (
	beginAdminState   = InputSG.New("startAdmin")
	adminCreateState  = InputSG.New("AdminCreateState")
	adminConnectState = InputSG.New("AdminConnectState")
	onCreateState     = InputSG.New("onCreateState")
	onConnectState    = InputSG.New("onConnectState")
)

func initAdminHandlers(manager *fsm.Manager, db *sql.DB) {
	//start buttons
	manager.Bind(&keyboards.BtnAdmin, fsm.DefaultState, onStartAdmin(db))
	//create
	manager.Bind(&keyboards.BtnCreate, beginAdminState, onCreating)
	initCreatingHandlers(manager, db)
	//connect
	manager.Bind(&keyboards.BtnConnect, beginAdminState, onConnecting)
	initConnectingHandlers(manager, db)
	//additional

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
	stmt := `SELECT COUNT(*) FROM Admins WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
