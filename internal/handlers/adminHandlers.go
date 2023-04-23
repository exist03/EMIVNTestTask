package handlers

import (
	"EMIVNTestTask/internal/users"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	"strconv"
)

func initAdminHandlers(command []string, db *sql.DB, id string) string {
	switch command[0] {
	case "create_card": //admin create_card [cardID] [bankInfo] [LimitInfo] - optional
		bankInfo := command[2]
		cardID, _ := strconv.Atoi(command[1])
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
	case "connect_card": //admin connect_card [cardID] [owner]
		cardID, _ := strconv.Atoi(command[1])
		owner := command[2]
		cardModel := mysql.CardModel{DB: db}
		cardModel.SetOwner(cardID, owner)
		return "Done"
	case "create_shogun": //admin create_shogun [Nickname] [TG username]
		nickname := command[1]
		username := command[2]
		shogun := users.Shogun{
			TelegramUsername: username,
			Nickname:         nickname,
		}
		shogunModel := mysql.ShogunModel{DB: db}
		err := shogunModel.Insert(shogun)
		if err != nil {
			return "Something went wrong"
		}
		return "Done"
	case "create_daimyo": //admin create_daimyo [Nickname] [TG username]
		nickname := command[1]
		username := command[2]
		daimyo := users.Daimyo{
			Owner:            "1",
			TelegramUsername: username,
			Nickname:         nickname,
		}
		daimyoModel := mysql.DaimyoModel{DB: db}
		err := daimyoModel.Insert(daimyo)
		if err != nil {
			return "Something went wrong"
		}
		return "Done"
	case "create_samurai": // admin create_samurai [Nickname] [TG username]
		nickname := command[1]
		username := command[2]
		daimyo := users.Samurai{
			Owner:            "1",
			TelegramUsername: username,
			Nickname:         nickname,
		}
		samuraiModel := mysql.SamuraiModel{DB: db}
		err := samuraiModel.Insert(daimyo)
		if err != nil {
			return "Something went wrong"
		}
		return "Done"
	case "create_collector":
		nickname := command[1]
		username := command[2]
		daimyo := users.Collector{
			TelegramUsername: username,
			Nickname:         nickname,
		}
		collectorModel := mysql.CollectorModel{DB: db}
		err := collectorModel.Insert(daimyo)
		if err != nil {
			return "Something went wrong"
		}
		return "Done"
	}
	return "_______"
}
