package tests

import (
	"banco_de_espana/repository"
	"log"
	"os"
	"testing"
	// "."
)

func TestMain(m *testing.M) {
	ensureTableExists()

	code := m.Run()
	clearTable()

	os.Exit(code)

}

func ensureTableExists() {
	if _, err := repository.ConnDB.Exec(tableCreationQuery()); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	repository.ConnDB.Exec("DELETE FROM clients")
	repository.ConnDB.Exec("DELETE FROM accounts")
}

func tableCreationQuery() string {
	return "select 1;"
}

// func tableCreationQuery() string {
// 	sql, err := ioutil.ReadFile("../infrastructure/schema_up.sql")
// 	println(string(sql))
// 	if err != nil {
// 		panic(err)
// 	}

// 	return string(sql)
// 	return "select 1;"
// }
