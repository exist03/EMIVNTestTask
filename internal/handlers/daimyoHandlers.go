package handlers

import (
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	"strconv"
)

func initDaimyoHandlers(command []string, db *sql.DB, id string) string {
	if command[0] == "subordinates" {
		samuraiModel := mysql.SamuraiModel{DB: db}
		res, _ := samuraiModel.GetList(id)
		return res
	} else if command[0] == "cards" {
		cardModel := mysql.CardModel{DB: db}
		res, _ := cardModel.GetList(id)
		return res
	} else if command[0] == "application" { //daimyo application [card id] [sum]
		daimyoModel := mysql.DaimyoModel{DB: db}
		res := daimyoModel.InsertApp(id, command[1], command[2])
		return res
	} else if command[0] == "set" { //daimyo set [card id] [sum]
		cardModel := mysql.CardModel{DB: db}
		cardID, err := strconv.Atoi(command[1])
		if err != nil {
			return "Something went wrong"
		}
		balance, err := strconv.ParseFloat(command[2], 64)
		if err != nil {
			return "Something went wrong"
		}
		res := cardModel.Update(cardID, balance)
		return res
	}
	return "Wrong message"
}
