package mysql

import (
	"EMIVNTestTask/internal/users"
	"database/sql"
)

type DaimyoModel struct {
	DB *sql.DB
}

func (m *DaimyoModel) Insert(daimyo users.Daimyo) error {
	stmt := `INSERT INTO Daimyo (Nickname, Owner, TelegramUsername)
   VALUES(?, ?, ?)`

	_, err := m.DB.Exec(stmt, daimyo.Nickname, daimyo.Owner.Nickname, daimyo.TelegramUsername)
	if err != nil {
		return err
	}
	return nil
}

func (m *DaimyoModel) InsertApp(daimyo users.Daimyo) error {
	stmt := `INSERT INTO Daimyo (Nickname, Owner, TelegramUsername)
   VALUES(?, ?, ?)`

	_, err := m.DB.Exec(stmt, daimyo.Nickname, daimyo.Owner.Nickname, daimyo.TelegramUsername)
	if err != nil {
		return err
	}
	return nil
}
