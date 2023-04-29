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

	_, err := m.DB.Exec(stmt, samurai.Nickname, samurai.Owner, samurai.TurnOver, samurai.TelegramUsername)
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

func (m *SamuraiModel) SetTurnover(id string, val float64) string {
	stmt := `UPDATE Samurais SET TurnOver = ? WHERE TelegramUsername = ?`
	_, err := m.DB.Exec(stmt, val, id)
	if err != nil {
		return "Something went wrong"
	}
	return "Done"
}

func (m *SamuraiModel) SetOwner(ID string, owner string) string {

	stmt := `UPDATE Samurais SET Owner=? WHERE TelegramUsername=?;`
	_, err := m.DB.Exec(stmt, owner, ID)
	if err != nil {
		return "Something went wrong"
	}
	return "Done"
}

func (m *SamuraiModel) Get(nickname string) string {
	stmt := `SELECT TurnOver, TelegramUsername, Owner, Nickname FROM Samurais WHERE TelegramUsername=?`
	row := m.DB.QueryRow(stmt, nickname)
	samurai := users.Samurai{}
	row.Scan(&samurai.TurnOver, &samurai.TelegramUsername, &samurai.Owner, &samurai.Nickname)
	return fmt.Sprintf("%s\n", samurai)
}
