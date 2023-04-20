package mysql

import (
	"EMIVNTestTask/internal/users"
	"database/sql"
	"fmt"
	"log"
)

type SamuraiModel struct {
	DB *sql.DB
}

func (m *SamuraiModel) Insert(samurai users.Samurai) error {
	stmt := `INSERT INTO Samurais (Nickname, Owner, TurnOver, TelegramUsername) VALUES(?, ?, ?, ?)`

	_, err := m.DB.Exec(stmt, samurai.Nickname, samurai.Owner.Nickname, samurai.TurnOver, samurai.TelegramUsername)
	if err != nil {
		return err
	}
	return nil
}

func (m *SamuraiModel) GetList(nickname string) (string, error) {
	stmt := `SELECT Nickname, TurnOver, TelegramUsername FROM Samurais WHERE Owner = ?`

	rows, err := m.DB.Query(stmt, nickname)
	if err != nil {
		log.Print(err)
		return "err_sql_query", err
	}
	defer rows.Close()

	var result string

	for rows.Next() {
		s := &users.Samurai{}
		err = rows.Scan(&s.Nickname, &s.TurnOver, &s.TelegramUsername)
		if err != nil {
			return "err_scan", err
		}
		result += fmt.Sprintf("%s", s)
	}

	if err = rows.Err(); err != nil {
		return "err3", err
	}
	return result, nil
}
