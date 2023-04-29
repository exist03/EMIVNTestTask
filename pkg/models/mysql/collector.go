package mysql

import (
	"EMIVNTestTask/internal/users"
	"database/sql"
	"fmt"
	"log"
)

type CollectorModel struct {
	DB *sql.DB
}

func (m *CollectorModel) Insert(collector users.Collector) error {
	stmt := `INSERT INTO Collectors (Nickname, TelegramUsername)
   VALUES(?, ?)`

	_, err := m.DB.Exec(stmt, collector.Nickname, collector.TelegramUsername)
	if err != nil {
		return err
	}
	return nil
}

func (m *CollectorModel) ShowApplications() string {
	stmt := `SELECT ID, Daimyo, SUM FROM Applications`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		log.Print(err)
		return "err_sql_query"
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		var (
			id     int
			sum    float64
			daimyo string
		)
		err = rows.Scan(&id, &daimyo, &sum)
		if err != nil {
			return "err_scan"
		}
		result += fmt.Sprintf("Daymo: %s\nCard id: %d\nSum: %.2f\n\n", daimyo, id, sum)
	}
	if err = rows.Err(); err != nil {
		return "err3"
	}
	return result
}

func (m *CollectorModel) ApplyApplication(cardID int, balance float64) string {
	stmtDelete := `DELETE FROM Applications WHERE ID=?;`
	stmtUpdate := `UPDATE Cards SET Balance = ? WHERE ID=?;`

	m.DB.Query(stmtUpdate, balance, cardID)
	m.DB.Query(stmtDelete, cardID)
	return "DONE"
}

func (m *CollectorModel) Get(nickname string) string {
	stmt := `SELECT TelegramUsername, Nickname FROM Collectors WHERE TelegramUsername=?`
	row := m.DB.QueryRow(stmt, nickname)
	collector := users.Collector{}
	row.Scan(&collector.TelegramUsername, &collector.Nickname)
	return fmt.Sprintf("%s\n", collector)
}
