package mysql

import "database/sql"

type ReportModel struct {
	DB *sql.DB
}

func (r *ReportModel) Samurai() string {
	var (
		startAmount  float64
		postupleniya float64
		spisaniya    float64
		finishAmount float64
	)

	stmt := ``
}
