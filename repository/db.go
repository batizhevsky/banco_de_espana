package repository

import (
	"database/sql"
)

var connDB = dbConnect()

func dbConnect() *sql.DB {
	connDB, err := sql.Open("sqlite3", "file:/tmp/bank.db?cache=shared&mode=rwc")
	checkErr(err)
	return connDB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
