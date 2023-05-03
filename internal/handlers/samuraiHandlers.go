package handlers

import (
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	"log"
	"strconv"
)

func initSamuraiHandlers(command []string, db *sql.DB, id string) string {
	if command[0] == "turnover" { //samurai turnover [value]
		samuraiModel := mysql.SamuraiModel{DB: db}
		value, err := strconv.ParseFloat(command[1], 64)
		if err != nil {
			log.Println(err, "___")
			return "Something went wrong"
		}
		res := samuraiModel.SetTurnover(id, value)
		return res
	}
	return "Incorrect message"
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
