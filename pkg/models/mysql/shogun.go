package mysql

import (
	"EMIVNTestTask/internal/users"
	"database/sql"
	"fmt"
)

type ShogunModel struct {
	DB *sql.DB
}

func (m *ShogunModel) Insert(shogun users.Shogun) error {
	stmt := `INSERT INTO Shogun (Nickname, TelegramUsername)
   VALUES(?, ?)`

	_, err := m.DB.Exec(stmt, shogun.Nickname, shogun.TelegramUsername)
	if err != nil {
		return err
	}
	return nil
}

func (m *ShogunModel) Get(nickname string) string {
	stmt := `SELECT TelegramUsername, Nickname FROM Shogun WHERE Nickname=?`
	row := m.DB.QueryRow(stmt, nickname)
	shogun := users.Shogun{}
	row.Scan(&shogun.TelegramUsername, &shogun.Nickname)
	result := fmt.Sprintf("%sDaimyos:\n", shogun)
	stmt1 := `SELECT Nickname FROM Daimyo WHERE Owner=?`
	rows, _ := m.DB.Query(stmt1, nickname)
	defer rows.Close()
	for rows.Next() {
		var temp string
		rows.Scan(&temp)
		result += fmt.Sprintf("%s ", temp)
	}
	return result
}
