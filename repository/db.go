package repository

import "database/sql"

func DBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./bank.db")
	checkErr(err)
	return db
}
