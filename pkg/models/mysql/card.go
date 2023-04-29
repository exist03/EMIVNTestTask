package mysql

import (
	"EMIVNTestTask/internal/users"
	"database/sql"
	"fmt"
	"log"
)

type CardModel struct {
	DB *sql.DB
}

func (m *CardModel) Insert(card users.Card) error {
	stmt := `INSERT INTO Cards (Owner, BankInfo, LimitInfo, Balance, ID) VALUES(?, ?, ?, ?, ?)`

	_, err := m.DB.Exec(stmt, card.Owner, card.BankInfo, card.LimitInfo, card.Balance, card.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *CardModel) GetList(nickname string) (string, error) {
	stmt := `SELECT Owner, ID, BankInfo, LimitInfo, Balance FROM Cards WHERE Owner = ?`

	rows, err := m.DB.Query(stmt, nickname)
	if err != nil {
		log.Print(err)
		return "err_sql_query", err
	}
	defer rows.Close()

	var result string

	for rows.Next() {
		s := &users.Card{}
		err = rows.Scan(&s.Owner, &s.ID, &s.BankInfo, &s.LimitInfo, &s.Balance)
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

func (m *CardModel) Update(id int, balance float64) string {
	stmt := `UPDATE Cards SET Balance=? WHERE ID=?;`

	_, err := m.DB.Exec(stmt, balance, id)
	if err != nil {
		return "Something went wrong"
	}
	return "Done"
}

func (m *CardModel) SetOwner(cardID string, owner string) string {
	stmt := `UPDATE Cards SET Owner=? WHERE ID=?;`

	_, err := m.DB.Exec(stmt, owner, cardID)
	if err != nil {
		return "Something went wrong"
	}
	return "Done"
}
