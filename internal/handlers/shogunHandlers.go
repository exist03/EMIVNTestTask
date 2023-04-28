package handlers

import (
	"EMIVNTestTask/internal/users"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	"strconv"
)

func initShogunHandlers(command []string, db *sql.DB, id string) string {
	switch command[0] {
	case "daimyos": //подчиненные дайме
		daimyoModel := mysql.DaimyoModel{DB: db}
		res, _ := daimyoModel.GetList(id)
		return res
	case "samurais": //shogun samurais [daimyo id]
		daimyoID := command[1]
		samuraiModel := mysql.SamuraiModel{DB: db}
		res, _ := samuraiModel.GetList(daimyoID)
		return res
	case "create": //shogun create [cardID] [bankInfo] [LimitInfo] - optional
		bankInfo := command[2]
		cardID := command[1]
		var limit float64
		if len(command) != 4 {
			limit = 2000000
		} else {
			limit, _ = strconv.ParseFloat(command[3], 64)
		}
		card := users.Card{
			ID:        cardID,
			Owner:     "1",
			BankInfo:  bankInfo,
			LimitInfo: limit,
			Balance:   limit,
		}
		cardModel := mysql.CardModel{DB: db}
		err := cardModel.Insert(card)
		if err != nil {
			return "Something went wrong"
		}
		return "Done"
	case "connect": //shogun connect [cardID] [owner]
		cardID := command[1]
		owner := command[2]
		cardModel := mysql.CardModel{DB: db}
		cardModel.SetOwner(cardID, owner)
		return "Done"
	}
	return "Wrong message"
}
