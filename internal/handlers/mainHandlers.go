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

		if sl[0] == "admin" {

		} else if sl[0] == "shogun" {
			return c.Send(initShogunHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		} else if sl[0] == "daimyo" {
			return c.Send(initDaimyoHandlers(sl[1:], db, strconv.Itoa(int(c.Sender().ID))))
		} else if sl[0] == "samumrai" {
			//return (c.Send(initSamuraiHandlers(sl[1:])))
		} else if sl[0] == "collector" {

		}
		return c.Send("Incorrect message")
	})
}
