package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // ...
)

var ConnDB = dbConnect()

func dbConnect() *sql.DB {
	ConnDB, err := sql.Open("sqlite3", "file:/tmp/bank.sqlite?cache=shared&mode=rwc")
	checkErr(err)
	return ConnDB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
