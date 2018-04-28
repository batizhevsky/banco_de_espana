package main_test

import (
	"log"
	"os"
	"testing"

	"."
)

func TestMain(m *testing.M) {
	ensureTableExists()

	code := m.Run()
	clearTable()

	os.Exit(code)

}

func ensureTableExists() {
	if _, err := repository.dbConnect.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}
