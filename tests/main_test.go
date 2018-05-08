package tests

import (
	"banco_de_espana/repository"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	// "."
)

func TestMain(m *testing.M) {
	repository.EnsureTableExists()
	repository.ClearTables()

	code := m.Run()
	repository.ClearTables()

	os.Exit(code)
}

func ensureTableExists() {
	file, err := ioutil.ReadFile("../infrastructure/schema_up.sql")
	if err != nil {
		panic(err)
	}

	requests := strings.Split(string(file), ";") // supposed it is not contain csv
	for _, request := range requests {
		if request == "" {
			continue
		}

		if _, err := repository.ConnDB.Exec(request); err != nil {
			log.Fatal(err)
		}
	}
}
