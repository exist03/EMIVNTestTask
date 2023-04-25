package handlers

import (
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	"strconv"
)

func initDaimyoHandlers(command []string, db *sql.DB, id string) string {
	switch command[0] {
	case "samurais": //daimyo samurais
		samuraiModel := mysql.SamuraiModel{DB: db}
		res, _ := samuraiModel.GetList(id)
		return res
	case "cards": //daimyo cards
		cardModel := mysql.CardModel{DB: db}
		res, _ := cardModel.GetList(id)
		return res
	case "application": //daimyo application [cardID] [value]
		daimyoModel := mysql.DaimyoModel{DB: db}
		res := daimyoModel.InsertApp(id, command[1], command[2])
		return res
	case "set": //daimyo set [cardID] [balance]
		if len(command) != 3 {
			return "Wrong message"
		}
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
