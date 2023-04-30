package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type ReportModel struct {
	DB *sql.DB
}

func (r *ReportModel) Samurai(id string, t time.Time) string {
	var (
		//startAmount  float64 //сумма на начало смены
		entrance float64 //поступления за смену
		offs     float64 // списания за смену
		//finishAmount float64 // сумма на конец смены
	)
	t = t.Add(time.Hour * 8)
	stmtEntrance := `SELECT SUM(Amount) FROM Transactions WHERE OperationType=1 AND SamuraiUsername=? AND Date BETWEEN ? AND ?`
	stmtOff := `SELECT SUM(Amount) FROM Transactions WHERE OperationType=0 AND SamuraiUsername=? AND Date BETWEEN ? AND ?`
	rowEntranse := r.DB.QueryRow(stmtEntrance, id, t, t.Add(time.Hour*24))
	rowOff := r.DB.QueryRow(stmtOff, id, t, t.Add(time.Hour*24))
	rowEntranse.Scan(&entrance)
	rowOff.Scan(&offs)
	//date cardid typeof operation
	stmtOperations := `SELECT Date, CardID, OperationType, Amount FROM Transactions WHERE SamuraiUsername=? AND Date BETWEEN ? AND ?`
	rows, err := r.DB.Query(stmtOperations, id, t, t.Add(time.Hour*24))
	if err != nil {
		log.Print(err)
		return "err_sql_query"
	}
	defer rows.Close()
	var (
		d             string
		cardID        string
		operationType bool
		amount        float64
	)
	var operations string
	for rows.Next() {
		err := rows.Scan(&d, &cardID, &operationType, &amount)
		if err != nil {
			log.Print(err)
		}
		operations += fmt.Sprintf("%19s | %10s | %v | %.2f\n", d, cardID, operationType, amount)
	}
	return fmt.Sprintf("Отчет за %v\nПоступления: %.2f\nСписания: %.2f\n%s\n", t.Format("2006-01-02"), entrance, offs, operations)
}
