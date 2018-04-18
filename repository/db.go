package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // ...
)

var connDB = dbConnect()

func dbConnect() *sql.DB {
	connDB, err := sql.Open("sqlite3", "file:/tmp/bank.sqlite?cache=shared&mode=rwc")
	checkErr(err)
	return connDB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
