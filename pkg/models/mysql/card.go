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
	stmt := `INSERT INTO Card (Owner, BankInfo, LimitInfo, Balance, ID) VALUES(?, ?, ?, ?, ?)`

	_, err := m.DB.Exec(stmt, card.Owner, card.BankInfo, card.LimitInfo, card.Balance, card.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *CardModel) CardLimit(nickname string) ([]string, error) {
	stmt := `SELECT ID, LimitInfo FROM Card WHERE Owner = ?`

	rows, err := m.DB.Query(stmt, nickname)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]string, 3)

	for rows.Next() {
		var id, limit string
		err = rows.Scan(&id, &limit)
		if err != nil {
			return nil, err
		}
		result = append(result, fmt.Sprintf("%s - %s", id, limit))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *CardModel) GetList(nickname string) ([]string, error) {
	stmt := `SELECT Owner, ID, BankInfo, LimitInfo, Balance FROM Card WHERE Owner = ?`

	rows, err := m.DB.Query(stmt, nickname)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]string, 3)

	for rows.Next() {
		s := &users.Card{}
		err = rows.Scan(&s.Owner, &s.ID, &s.BankInfo, &s.LimitInfo, &s.Balance)
		if err != nil {
			return nil, err
		}
		result = append(result, fmt.Sprintf("%s", s))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *CardModel) Update(id interface{}, balance float64) string {
	stmt := `UPDATE Card SET Balance=? WHERE ID=?;`

	_, err := m.DB.Exec(stmt, balance, id)
	if err != nil {
		return "Something went wrong"
	}
	return "Done"
}

func (m *CardModel) SetOwner(cardID interface{}, owner string) string {
	stmt := `UPDATE Card SET Owner=? WHERE ID=?;`
	_, err := m.DB.Exec(stmt, owner, cardID)
	if err != nil {
		return "Something went wrong"
	}
	return "Done"
}
