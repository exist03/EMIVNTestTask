package handlers

import (
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
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
	case "":

	}
	return "Wrong message"
}
