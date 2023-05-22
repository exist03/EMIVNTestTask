package mysql

import (
	"EMIVNTestTask/internal/users"
	"database/sql"
	"fmt"
	"log"
	"time"
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

func (m *ShogunModel) Get(nickname string) string {
	stmt := `SELECT TelegramUsername, Nickname FROM Shogun WHERE TelegramUsername=?`
	row := m.DB.QueryRow(stmt, nickname)
	shogun := users.Shogun{}
	row.Scan(&shogun.TelegramUsername, &shogun.Nickname)
	result := fmt.Sprintf("%sDaimyos:\n", shogun)
	stmt1 := `SELECT TelegramUsername FROM Daimyo WHERE Owner=?`
	rows, _ := m.DB.Query(stmt1, nickname)
	defer rows.Close()
	for rows.Next() {
		var temp string
		rows.Scan(&temp)
		result += fmt.Sprintf("%s ", temp)
	}
	return result
}

type daimyoStruct struct {
	CardSum         float64
	AdditionSum     float64
	TurnoverSamurai float64
	RemainingFunds  float64
}

func (m *ShogunModel) GetReportShift(shogun interface{}) (map[string]daimyoStruct, error) {
	date := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	today := fmt.Sprintf("%s 8:00", time.Now().Format("2006-01-02"))
	yesterday := fmt.Sprintf("%s 8:00", time.Now().Add(-24*time.Hour).Format("2006-01-02"))
	myMap := make(map[string]daimyoStruct)
	stmtCardSum := `SELECT TelegramUsername, SUM(CBD.Amount) FROM Daimyo
JOIN  Card ON Daimyo.TelegramUsername = Card.Owner
JOIN CardBalanceDaily CBD ON Card.ID = CBD.CardID
WHERE CBD.Time=? AND Daimyo.Owner=? GROUP BY TelegramUsername;`
	stmtAdditionSum := `SELECT TelegramUsername, SUM(T.Amount) FROM Daimyo
JOIN Card ON Daimyo.TelegramUsername = Card.Owner
JOIN Transaction T on Card.ID = T.CardID WHERE T.OperationType=true AND Daimyo.Owner=? AND T.Date BETWEEN ? AND ? GROUP BY TelegramUsername;`
	stmtTurnoverSamuraiEnd := `SELECT Daimyo.TelegramUsername, SUM(TE.Amount) FROM Daimyo
JOIN Samurai S ON Daimyo.TelegramUsername = S.Owner
JOIN TurnoverEnd TE ON TE.SamuraiUsername = S.TelegramUsername
WHERE Daimyo.Owner = ? AND TE.Date = ? GROUP BY Daimyo.TelegramUsername;`
	stmtTurnoverSamuraiBegin := `SELECT Daimyo.TelegramUsername, SUM(TB.Amount) FROM Daimyo
JOIN Samurai S ON Daimyo.TelegramUsername = S.Owner
JOIN TurnoverBegin TB ON TB.SamuraiUsername = S.TelegramUsername
WHERE Daimyo.Owner = ? AND TB.Date = ? GROUP BY Daimyo.TelegramUsername;`

	rows, err := m.DB.Query(stmtCardSum, date, shogun)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var daimyo string
		var amount float64
		err = rows.Scan(&daimyo, &amount)
		if err != nil {
			return nil, err
		}
		s := daimyoStruct{
			CardSum: amount,
		}
		myMap[daimyo] = s
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	rowsAdditionalSum, err := m.DB.Query(stmtAdditionSum, shogun, yesterday, today)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsAdditionalSum.Close()
	for rowsAdditionalSum.Next() {
		var daimyo string
		var amount float64
		err = rowsAdditionalSum.Scan(&daimyo, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[daimyo]; ok {
			str.AdditionSum = amount
			myMap[daimyo] = str
		} else {
			s := daimyoStruct{
				AdditionSum: amount,
			}
			myMap[daimyo] = s
		}
	}
	if err = rowsAdditionalSum.Err(); err != nil {
		return nil, err
	}

	rowsTurnoverSamuraiEnd, err := m.DB.Query(stmtTurnoverSamuraiEnd, shogun, date)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsTurnoverSamuraiEnd.Close()

	for rowsTurnoverSamuraiEnd.Next() {
		var daimyo string
		var amount float64
		err = rowsTurnoverSamuraiEnd.Scan(&daimyo, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[daimyo]; ok {
			str.TurnoverSamurai = amount
			myMap[daimyo] = str
		} else {
			s := daimyoStruct{
				TurnoverSamurai: amount,
			}
			myMap[daimyo] = s
		}
	}
	if err = rowsTurnoverSamuraiEnd.Err(); err != nil {
		return nil, err
	}

	rowsTurnoverSamuraiBegin, err := m.DB.Query(stmtTurnoverSamuraiBegin, shogun, date)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsTurnoverSamuraiBegin.Close()

	for rowsTurnoverSamuraiBegin.Next() {
		var daimyo string
		var amount float64
		err = rowsTurnoverSamuraiBegin.Scan(&daimyo, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[daimyo]; ok {
			str.TurnoverSamurai -= amount
			myMap[daimyo] = str
		} else {
			s := daimyoStruct{
				TurnoverSamurai: amount,
			}
			myMap[daimyo] = s
		}
	}
	if err = rowsTurnoverSamuraiBegin.Err(); err != nil {
		return nil, err
	}
	return myMap, nil
}

//
//func (m *ShogunModel) GetReportPeriod(shogun interface{}, dateBegin, dateEnd interface{}) (map[string]SamuraiStruct, error) {
//	myMap := make(map[string]SamuraiStruct)
//	stmtSamuraiSber := `Select SamuraiUsername, TurnoverEnd.Amount
//	from TurnoverEnd
//	join Samurai S on TurnoverEnd.SamuraiUsername = S.TelegramUsername
//	where S.owner=? AND Bank="сбер" AND Date=?`
//
//	stmtSamuraiTink := `Select SamuraiUsername, TurnoverEnd.Amount
//	from TurnoverEnd
//	join Samurai S on TurnoverEnd.SamuraiUsername = S.TelegramUsername
//	where S.owner=? AND Bank="тинькофф" AND Date=?`
//
//	rows, err := m.DB.Query(stmtSamuraiSber, daimyo, dateBegin)
//	if err != nil {
//		log.Print(err)
//		return nil, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var samurai string
//		var amount float64
//		err = rows.Scan(&samurai, &amount)
//		log.Printf("%s sber begin = %.0f\n", samurai, amount)
//		if err != nil {
//			return nil, err
//		}
//		s := SamuraiStruct{
//			Sber: amount,
//		}
//		myMap[samurai] = s
//	}
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	rowsSberEnd, err := m.DB.Query(stmtSamuraiSber, daimyo, dateEnd)
//	if err != nil {
//		log.Print(err)
//		return nil, err
//	}
//	defer rowsSberEnd.Close()
//
//	for rowsSberEnd.Next() {
//		var samurai string
//		var amount float64
//		err = rowsSberEnd.Scan(&samurai, &amount)
//		log.Printf("%s sber end = %.0f\n", samurai, amount)
//		if err != nil {
//			return nil, err
//		}
//		if str, ok := myMap[samurai]; ok {
//			str.SberCheck = amount
//			myMap[samurai] = str
//		} else {
//			s := SamuraiStruct{
//				SberCheck: amount,
//			}
//			myMap[samurai] = s
//		}
//	}
//
//	rowsTink, err := m.DB.Query(stmtSamuraiTink, daimyo, dateBegin)
//	if err != nil {
//		log.Print(err)
//		return nil, err
//	}
//	defer rowsTink.Close()
//
//	for rowsTink.Next() {
//		var samurai string
//		var amount float64
//		err = rowsTink.Scan(&samurai, &amount)
//		log.Printf("%s sber begin = %.0f\n", samurai, amount)
//		if err != nil {
//			return nil, err
//		}
//		if str, ok := myMap[samurai]; ok {
//			str.Sber += amount
//			myMap[samurai] = str
//		} else {
//			s := SamuraiStruct{
//				SberCheck: amount,
//			}
//			myMap[samurai] = s
//		}
//	}
//
//	rowsTinkEnd, err := m.DB.Query(stmtSamuraiTink, daimyo, dateEnd)
//	if err != nil {
//		log.Print(err)
//		return nil, err
//	}
//	defer rowsTinkEnd.Close()
//
//	for rowsTinkEnd.Next() {
//		var samurai string
//		var amount float64
//		err = rowsTinkEnd.Scan(&samurai, &amount)
//		log.Printf("%s sber end = %.0f\n", samurai, amount)
//		if err != nil {
//			return nil, err
//		}
//		if str, ok := myMap[samurai]; ok {
//			str.SberCheck += amount
//			myMap[samurai] = str
//		} else {
//			s := SamuraiStruct{
//				SberCheck: amount,
//			}
//			myMap[samurai] = s
//		}
//	}
//
//	return myMap, nil
//}
