package mysql

import (
	"EMIVNTestTask/internal/users"
	"database/sql"
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
