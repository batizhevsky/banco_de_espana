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
	return `
	Create Table if not exists clients(
		id integer primary key asc,
		name text,
		email text,
		phone integer
	);
	
	Create Table if not exists accounts(
		id integer primary key asc,
		client_id integer,
		balance integer
	);
	
	Create Table if not exists transactions(
		id integer primary key asc,
		account_id integer,
		amount integer,
		subject text,
		timestamp integer 
	);
	`
}
