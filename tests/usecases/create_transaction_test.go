package usecases_test

import (
	"testing"

	"banco_de_espana/entities"
	"banco_de_espana/usecases"
)

func TestCreateTransaction(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", "+18005687625")
	acc, _ := entities.NewAccount(cl, 100.00)

	usecases.CreateTransaction(acc, 70.00)

	if acc.Balance != 30.0 {
		t.Error("Expected", 30.0, "Got", acc.Balance)
	}
}
