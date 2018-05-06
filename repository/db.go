package repository

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// ConnDB stores a connection
var ConnDB = dbConnect()

const dbName = "banko_de_espana"

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

func database() string {
	if isTestEnv() {
		return fmt.Sprintf("%s_test", dbName)
	}
	return dbName
}

func isTestEnv() bool {
	return flag.Lookup("test.v") != nil
}
