package mysql

import (
	"EMIVNTestTask/internal/users"
	"database/sql"
	"fmt"
	"log"
)

type DaimyoModel struct {
	DB *sql.DB
}

func (m *DaimyoModel) Insert(daimyo users.Daimyo) error {
	stmt := `INSERT INTO Daimyo (Nickname, Owner, TelegramUsername)
   VALUES(?, ?, ?)`

	_, err := m.DB.Exec(stmt, daimyo.Nickname, daimyo.Owner, daimyo.TelegramUsername)
	if err != nil {
		return err
	}
	return nil
}

func (m *DaimyoModel) InsertApp(creater, cardID, sum string) string {
	stmt := `INSERT INTO Applications (Daimyo, ID, Sum) VALUES(?, ?, ?)`
	_, err := m.DB.Exec(stmt, creater, cardID, sum)
	if err != nil {
		return "Something went wrong"
	}
	return "New application created"
}

func (m *DaimyoModel) GetList(owner string) (string, error) {
	stmt := `SELECT TelegramUsername, Nickname FROM Daimyo WHERE Owner = ?`

	rows, err := m.DB.Query(stmt, owner)
	if err != nil {
		log.Print(err)
		return "err_sql_query", err
	}
	defer rows.Close()

	var result string

	for rows.Next() {
		//s := &users.Daimyo{}
		var telegramUsername string
		var nickname string
		err = rows.Scan(&telegramUsername, &nickname)
		if err != nil {
			return "err_scan", err
		}
		result += fmt.Sprintf("TG Username: %s\nNickname: %s", telegramUsername, nickname)
	}

	if err = rows.Err(); err != nil {
		return "err3", err
	}
	return result, nil
}

func (m *DaimyoModel) SetOwner(ID string, owner string) string {
	stmt := `UPDATE Daimyo SET Owner=? WHERE TelegramUsername=?;`
	_, err := m.DB.Exec(stmt, owner, ID)
	if err != nil {
		return "Something went wrong"
	}
	return "Done"
}

func (m *DaimyoModel) Get(username string) string {
	stmt := `SELECT Owner, TelegramUsername, Nickname FROM Daimyo WHERE TelegramUsername=?`
	row := m.DB.QueryRow(stmt, username)
	daimyo := users.Daimyo{}
	row.Scan(&daimyo.Owner, &daimyo.TelegramUsername, &daimyo.Nickname)
	result := fmt.Sprintf("%sSamurais:\n", daimyo)
	stmt1 := `SELECT Nickname FROM Samurais WHERE Owner=?`
	rows, _ := m.DB.Query(stmt1, username)
	defer rows.Close()
	for rows.Next() {
		var temp string
		rows.Scan(&temp)
		result += fmt.Sprintf("%s ", temp)
	}
	return result
}
