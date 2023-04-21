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

func (m *DaimyoModel) InsertApp(creater, cardID, sum string) string {
	stmt := `INSERT INTO Applications (Daimyo, ID, Sum) VALUES(?, ?, ?)`
	_, err := m.DB.Exec(stmt, creater, cardID, sum)
	if err != nil {
		return "Something went wrong"
	}
	return "New application created"
}
