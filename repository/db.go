package repository

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// ConnDB stores a connection
var ConnDB = dbConnect()

const dbName = "banko_de_espana"

func EnsureTableExists() {
	absPath, _ := filepath.Abs("../infrastructure/schema_up.sql")
	println(absPath)
	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	requests := strings.Split(string(file), ";") // supposed it is not contain csv
	for _, request := range requests {
		if request == "" {
			continue
		}

		if _, err := ConnDB.Exec(request); err != nil {
			log.Fatal(err)
		}
	}
}

func ClearTables() {
	ConnDB.Exec("DELETE FROM clients")
	ConnDB.Exec("DELETE FROM accounts")
	ConnDB.Exec("DELETE FROM transactions")
}

func dbConnect() *sql.DB {
	ConnDB, err := sql.Open("mysql", connection())
	checkErr(err)

	err = ConnDB.Ping()
	checkErr(err)

	return ConnDB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func connection() string {
	return fmt.Sprintf("root:dummy@/%s?charset=utf8", databaseName())
}

func databaseName() string {
	if isTestEnv() {
		return fmt.Sprintf("%s_test", dbName)
	}
	return dbName
}

func isTestEnv() bool {
	return flag.Lookup("test.v") != nil
}
