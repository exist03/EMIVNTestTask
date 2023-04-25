package handlers

import (
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	"strconv"
)

func initSamuraiHandlers(command []string, db *sql.DB, id string) string {
	if command[0] == "turnover" { //samurai turnover [value]
		samuraiModel := mysql.SamuraiModel{DB: db}
		value, err := strconv.ParseFloat(command[1], 64)
		if err != nil {
			return "Something went wrong"
		}
		res := samuraiModel.SetTurnover(id, value)
		return res
	}
	return "Incorrect message"
}
