package usecases

import (
	"testing"

	"banco_de_espana/entities"
)

func TestCreateTransaction(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	acc, _ := entities.NewAccount(cl, 100.00)

	CreateTransaction(acc, 70.00)

	if acc.Balance != 30.0 {
		t.Error("Expected", 30.0, "Got", acc.Balance)
	}
}
