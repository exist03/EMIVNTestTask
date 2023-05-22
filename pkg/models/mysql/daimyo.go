package mysql

import (
	"EMIVNTestTask/internal/users"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type DaimyoModel struct {
	DB *sql.DB
}

func (m *DaimyoModel) Insert(daimyo users.Daimyo) error {
	stmt := `INSERT INTO Daimyo (Nickname, Owner, TelegramUsername)
   VALUES(?, ?, ?)`

	_, err := m.DB.Exec(stmt, daimyo.Nickname, daimyo.Owner, daimyo.TelegramUsername)
	if err != nil {
		return err
	}
	return nil
}

func (m *DaimyoModel) InsertApp(creater string, cardID interface{}, sum float64) string {
	stmt := `INSERT INTO Application (Daimyo, ID, Sum) VALUES(?, ?, ?)`
	_, err := m.DB.Exec(stmt, creater, cardID, sum)
	if err != nil {
		log.Println(err)
		return "Something went wrong"
	}
	return "New application created"
}

func (m *DaimyoModel) GetList(owner string) (string, error) {
	stmt := `SELECT TelegramUsername, Nickname, Owner FROM Daimyo WHERE Owner = ?`

	rows, err := m.DB.Query(stmt, owner)
	if err != nil {
		log.Print(err)
		return "err_sql_query", err
	}
	defer rows.Close()

	var result string

	for rows.Next() {
		d := &users.Daimyo{}
		err = rows.Scan(&d.TelegramUsername, &d.Nickname, &d.Owner)
		if err != nil {
			return "err_scan", err
		}
		//result += fmt.Sprintf("TG Username: %s\nNickname: %s\n\n", telegramUsername, nickname)
		result += fmt.Sprintf("%s", d)
	}

	if err = rows.Err(); err != nil {
		return "err3", err
	}
	return result, nil
}

func (m *DaimyoModel) SetOwner(ID interface{}, owner string) string {
	stmt := `UPDATE Daimyo SET Owner=? WHERE TelegramUsername=?;`
	_, err := m.DB.Exec(stmt, owner, ID)
	if err != nil {
		return "Something went wrong"
	}
	return "Done"
}

func (m *DaimyoModel) Get(username string) string {
	stmt := `SELECT Owner, TelegramUsername, Nickname FROM Daimyo WHERE TelegramUsername=?`
	row := m.DB.QueryRow(stmt, username)
	daimyo := users.Daimyo{}
	row.Scan(&daimyo.Owner, &daimyo.TelegramUsername, &daimyo.Nickname)
	result := fmt.Sprintf("%sSamurais:\n", daimyo)
	stmt1 := `SELECT TelegramUsername FROM Samurai WHERE Owner=?`
	rows, _ := m.DB.Query(stmt1, username)
	defer rows.Close()
	for rows.Next() {
		var temp string
		rows.Scan(&temp)
		result += fmt.Sprintf("%s ", temp)
	}
	return result
}

type SamuraiStruct struct {
	Tink      float64
	Sber      float64
	TinkCheck float64
	SberCheck float64
}

// shit happends...
func (m *DaimyoModel) GetReportShift(daimyo interface{}) (map[string]SamuraiStruct, error) {
	date := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	myMap := make(map[string]SamuraiStruct)
	stmtSamuraiSber := `Select SamuraiUsername, TurnoverEnd.Amount 
	from TurnoverEnd
	join Samurai S on TurnoverEnd.SamuraiUsername = S.TelegramUsername
	where S.owner=? AND Bank="сбер" AND Date=?`

	stmtSamuraiTink := `Select SamuraiUsername, TurnoverEnd.Amount
	from TurnoverEnd
	join Samurai S on TurnoverEnd.SamuraiUsername = S.TelegramUsername
	where S.owner=? AND Bank="тинькофф" AND Date=?`

	stmtColtrolSber := `Select SamuraiUsername, TurnoverController.Amount
	from TurnoverController
	join Samurai S on TurnoverController.SamuraiUsername = S.TelegramUsername
	where S.owner=? AND Bank="сбер" AND Date=?`

	stmtColtrolTink := `Select SamuraiUsername, TurnoverController.Amount
	from TurnoverController
	join Samurai S on TurnoverController.SamuraiUsername = S.TelegramUsername
	where S.owner=? AND Bank="тинькофф" AND Date=?`

	rows, err := m.DB.Query(stmtSamuraiSber, daimyo, date)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var samurai string
		var amount float64
		err = rows.Scan(&samurai, &amount)
		if err != nil {
			return nil, err
		}
		s := SamuraiStruct{
			Sber: amount,
		}
		myMap[samurai] = s
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	rowsTink, err := m.DB.Query(stmtSamuraiTink, daimyo, date)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsTink.Close()

	for rowsTink.Next() {
		var samurai string
		var amount float64
		err = rowsTink.Scan(&samurai, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.Tink = amount
			myMap[samurai] = str
		} else {
			s := SamuraiStruct{
				Tink: amount,
			}
			myMap[samurai] = s
		}
	}
	if err = rowsTink.Err(); err != nil {
		return nil, err
	}

	rowsTinkControls, err := m.DB.Query(stmtColtrolTink, daimyo, date)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsTinkControls.Close()

	for rowsTinkControls.Next() {
		var samurai string
		var amount float64
		err = rowsTinkControls.Scan(&samurai, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.TinkCheck = amount
			myMap[samurai] = str
		} else {
			s := SamuraiStruct{
				TinkCheck: amount,
			}
			myMap[samurai] = s
		}
	}
	if err = rowsTinkControls.Err(); err != nil {
		return nil, err
	}

	rowsSberControls, err := m.DB.Query(stmtColtrolSber, daimyo, date)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsSberControls.Close()

	for rowsSberControls.Next() {
		var samurai string
		var amount float64
		err = rowsSberControls.Scan(&samurai, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.SberCheck = amount
			myMap[samurai] = str
		} else {
			s := SamuraiStruct{
				SberCheck: amount,
			}
			myMap[samurai] = s
		}
	}
	if err = rowsSberControls.Err(); err != nil {
		return nil, err
	}

	return myMap, nil
}

func (m *DaimyoModel) GetReportPeriod(daimyo interface{}, dateBegin, dateEnd interface{}) (map[string]SamuraiStruct, error) {
	myMap := make(map[string]SamuraiStruct)
	stmtSamuraiSber := `Select SamuraiUsername, TurnoverEnd.Amount 
	from TurnoverEnd
	join Samurai S on TurnoverEnd.SamuraiUsername = S.TelegramUsername
	where S.owner=? AND Bank="сбер" AND Date=?`

	stmtSamuraiTink := `Select SamuraiUsername, TurnoverEnd.Amount
	from TurnoverEnd
	join Samurai S on TurnoverEnd.SamuraiUsername = S.TelegramUsername
	where S.owner=? AND Bank="тинькофф" AND Date=?`

	rows, err := m.DB.Query(stmtSamuraiSber, daimyo, dateBegin)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var samurai string
		var amount float64
		err = rows.Scan(&samurai, &amount)
		log.Printf("%s sber begin = %.0f\n", samurai, amount)
		if err != nil {
			return nil, err
		}
		s := SamuraiStruct{
			Sber: amount,
		}
		myMap[samurai] = s
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	rowsSberEnd, err := m.DB.Query(stmtSamuraiSber, daimyo, dateEnd)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsSberEnd.Close()

	for rowsSberEnd.Next() {
		var samurai string
		var amount float64
		err = rowsSberEnd.Scan(&samurai, &amount)
		log.Printf("%s sber end = %.0f\n", samurai, amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.SberCheck = amount
			myMap[samurai] = str
		} else {
			s := SamuraiStruct{
				SberCheck: amount,
			}
			myMap[samurai] = s
		}
	}

	rowsTink, err := m.DB.Query(stmtSamuraiTink, daimyo, dateBegin)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsTink.Close()

	for rowsTink.Next() {
		var samurai string
		var amount float64
		err = rowsTink.Scan(&samurai, &amount)
		log.Printf("%s sber begin = %.0f\n", samurai, amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.Sber += amount
			myMap[samurai] = str
		} else {
			s := SamuraiStruct{
				SberCheck: amount,
			}
			myMap[samurai] = s
		}
	}

	rowsTinkEnd, err := m.DB.Query(stmtSamuraiTink, daimyo, dateEnd)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsTinkEnd.Close()

	for rowsTinkEnd.Next() {
		var samurai string
		var amount float64
		err = rowsTinkEnd.Scan(&samurai, &amount)
		log.Printf("%s sber end = %.0f\n", samurai, amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.SberCheck += amount
			myMap[samurai] = str
		} else {
			s := SamuraiStruct{
				SberCheck: amount,
			}
			myMap[samurai] = s
		}
	}

	return myMap, nil
}
