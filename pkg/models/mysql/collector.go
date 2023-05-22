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
	stmt := `INSERT INTO Collector (Nickname, TelegramUsername)
   VALUES(?, ?)`

	_, err := m.DB.Exec(stmt, collector.Nickname, collector.TelegramUsername)
	if err != nil {
		return err
	}
	return nil
}

func (m *CollectorModel) ShowApplications() string {
	stmt := `SELECT ID, Daimyo, SUM FROM Application`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		log.Print(err)
		return "err_sql_query"
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		var (
			id     string
			sum    float64
			daimyo string
		)
		err = rows.Scan(&id, &daimyo, &sum)
		if err != nil {
			return "err_scan"
		}
		result += fmt.Sprintf("Daymo: %s\nCard id: %s\nSum: %.2f\n\n", daimyo, id, sum)
	}
	if err = rows.Err(); err != nil {
		return "err3"
	}
	return result
}

func (m *CollectorModel) ApplyApplication(cardID interface{}, balance float64) string {
	stmtDelete := `DELETE FROM Application WHERE ID=?;`
	stmtUpdate := `UPDATE Card SET Balance = Balance + ? WHERE ID=?;`
	_, err := m.DB.Query(stmtUpdate, balance, cardID)
	if err != nil {
		log.Println(err)
		return "update err"
	}
	m.DB.Query(stmtDelete, cardID)
	return "DONE"
}

func (m *CollectorModel) Get(nickname string) string {
	stmt := `SELECT TelegramUsername, Nickname FROM Collector WHERE TelegramUsername=?`
	row := m.DB.QueryRow(stmt, nickname)
	collector := users.Collector{}
	row.Scan(&collector.TelegramUsername, &collector.Nickname)
	return fmt.Sprintf("%s\n", collector)
}
