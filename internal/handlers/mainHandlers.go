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
		switch sl[0] {
		case "admin":

		case "shogun":
			return c.Send(initShogunHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		case "daimyo":
			return c.Send(initDaimyoHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		case "samurai":
			return c.Send(initSamuraiHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		case "collector":
			return c.Send(initCollectorHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		}
		return c.Send("Incorrect message")
	})
}
