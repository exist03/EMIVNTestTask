package handlers

import (
	"database/sql"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func InitHanders(b *tele.Bot, db *sql.DB) {
	b.Handle(tele.OnText, func(c tele.Context) error {
		sl := strings.Split(c.Text(), " ")
		senderID := strconv.Itoa(int(c.Sender().ID))
		switch sl[0] {
		case "admin":
			if !validAdmin(db, senderID) {
				return c.Send("You don`t have rights")
			}
			return c.Send(initAdminHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		case "shogun":
			if !validShogun(db, senderID) {
				return c.Send("You don`t have rights")
			}
			return c.Send(initShogunHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		case "daimyo":
			if !validDaimyo(db, senderID) {
				return c.Send("You don`t have rights")
			}
			return c.Send(initDaimyoHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		case "samurai":
			if !validSamurai(db, senderID) {
				return c.Send("You don`t have rights")
			}
			return c.Send(initSamuraiHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		case "collector":
			if !validCollector(db, senderID) {
				return c.Send("You don`t have rights")
			}
			return c.Send(initCollectorHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		}
		return c.Send("Incorrect message")
	})
}

func validAdmin(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Admins WHERE Nickname=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func validShogun(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Shogun WHERE Nickname=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func validDaimyo(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Daimyo WHERE Nickname=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func validSamurai(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Samurais WHERE Nickname=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func validCollector(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Collectors WHERE Nickname=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
