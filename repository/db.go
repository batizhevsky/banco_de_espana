package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

var ConnDB = dbConnect()

func dbConnect() *sql.DB {
	ConnDB, err := sql.Open("mysql", "root:dummy@/banko_de_espana?charset=utf8")
	checkErr(err)
	return ConnDB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
