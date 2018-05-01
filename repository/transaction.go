package repository

import (
	"banco_de_espana/entities"
	"time"
)

// CreateTransaction ...
func CreateTransaction(tr *entities.Transaction) {
	stmt, err := ConnDB.Prepare("INSERT INTO transactions(account_id, amount, subject, timestamp) values(?, ?, ?, ?)")

	checkErr(err)

	res, err := stmt.Exec(tr.Account.ID, tr.Amount, tr.Subject, tr.DateTime.Unix())

	checkErr(err)

	id, err := res.LastInsertId()

	checkErr(err)

	tr.ID = id
}

// GetTransaction ...
func GetTransaction(id int64) *entities.Transaction {
	rows, err := ConnDB.Query("Select * FROM transactions where id = ? limit 1", id)

	checkErr(err)

	tr := entities.Transaction{}
	var accID int64
	var timestamp int64

	if rows.Next() {
		err = rows.Scan(&tr.ID, &accID, &tr.Amount, &tr.Subject, &timestamp)
		checkErr(err)

		tr.DateTime = time.Unix(timestamp, 0)

		return &tr
	}

	return nil
}
