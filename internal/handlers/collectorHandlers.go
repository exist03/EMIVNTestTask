package handlers

import (
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

var (
	BeginCollectorState = InputSG.New("startCollector")
)

func initCollectorHandlers(command []string, db *sql.DB, id string) string {
	switch command[0] {
	case "show": //collector show
		CollectorModel := mysql.CollectorModel{DB: db}
		res := CollectorModel.ShowApplications()
		return res
	case "apply": //collector apply [cardID] [value]
		CollectorModel := mysql.CollectorModel{DB: db}
		cardID, err := strconv.Atoi(command[1])
		if err != nil {
			return "Something went wrong"
		}
		balance, err := strconv.ParseFloat(command[2], 64)
		if err != nil {
			return "Something went wrong"
		}
		res := CollectorModel.ApplyApplication(cardID, balance) //collector apply [card id] [balance]
		return res
	}
	return "Wrong message"
}

func onStartCollector(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		if !validCollector(db, c.Sender().Username) {
			return c.Send("У вас нет прав")
		}
		state.Set(BeginCollectorState)
		return c.Send("Выберите действие") //keyboards.collectorKB
	}
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
