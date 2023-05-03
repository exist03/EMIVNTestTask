package handlers

import (
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	"strconv"
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
