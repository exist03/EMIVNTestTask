package mysql

import (
	"database/sql"
	"fmt"
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
	return fmt.Sprintf("Поступления: %.2f\nСписания: %.2f\n", entrance, offs)
}
