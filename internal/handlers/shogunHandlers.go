package handlers

import "database/sql"

func initShogunHandlers(command []string, db *sql.DB, id string) string {
	if command[0] == "show" {

	} else if command[0] == "describe" {
		return "decribe"
	}
	return "Wrong message"
}
