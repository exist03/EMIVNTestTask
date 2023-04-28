package handlers

import (
	"database/sql"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

const (
	adminCommands     = "- admin create_card [cardID] [bankInfo] {LimitInfo}\n- admin connect_card [cardID] [owner]\n- admin create_shogun [Nickname] [TG username]\n- admin create_daimyo [Nickname] [TG username]\n- admin create_samurai [Nickname] [TG username]\n- admin create_collector [Nickname] [TG username]\n- admin set_daimyo_owner [daimyo nickname] [shogun nickname]\n- admin set_samurai_owner [samurai nickname] [daimyo nickname]\n- admin get_shogun_info [shogunID]\n- admin get_daimyo_info [daimyoID]\n- admin get_samurai_info [samuraiID]\n- admin get_collector_info [collectorID]"
	shogunCommands    = "- shogun daimyos - Просмотр подчиненных(Даймё)\n- shogun samurais [daimyo nickname] - Просмотр подчиненных самураев у конкретного даймё\n- shogun create [cardID] [bankInfo] {LimitInfo}\n- shogun connect [cardID] [owner] - Привезка карты к даймё по нику"
	daimyoCommands    = "- daimyo samurais\n- daimyo set [cardID] [balance] - Остаток на карте\n- daimyo cards\n- daimyo application [cardID] [value] - Создание заявки на пополнение карты[cardID] до суммы [value]"
	samuraiCommands   = "- samurai turnover [value] - оборот за смену"
	collectorCommands = "- collector show - Показать запросы на пополнение\n- collector apply [cardID] [value] - Выполнить запрос на пополнение"
)

func InitHanders(b *tele.Bot, db *sql.DB) {
	beginHandlers(b)
	b.Handle(tele.OnText, func(c tele.Context) error {
		sl := strings.Split(c.Text(), " ")
		senderID := c.Sender().Username
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

func beginHandlers(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Введите /[роль] для подробностей")
	})
	b.Handle("/admin", func(c tele.Context) error {
		return c.Send(adminCommands)
	})
	b.Handle("/shogun", func(c tele.Context) error {
		return c.Send(shogunCommands)
	})
	b.Handle("/daimyo", func(c tele.Context) error {
		return c.Send(daimyoCommands)
	})
	b.Handle("/samurai", func(c tele.Context) error {
		return c.Send(samuraiCommands)
	})
	b.Handle("/collector", func(c tele.Context) error {
		return c.Send(collectorCommands)
	})
}

func validAdmin(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Admins WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func validShogun(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Shogun WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func validDaimyo(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Daimyo WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func validSamurai(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Samurais WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
func validCollector(db *sql.DB, senderID string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM Collectors WHERE TelegramUsername=?`
	row := db.QueryRow(stmt, senderID)
	row.Scan(&temp)
	if temp == 0 {
		return false
	}
	return true
}
