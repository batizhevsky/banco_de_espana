package tests

import (
	"testing"

	"banco_de_espana/entities"
	"banco_de_espana/usecases"
)

func TestCreateTransfer(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	fromAcc, _ := entities.NewAccount(cl, 100.00)
	toAcc, _ := entities.NewAccount(cl, 20.00)

	usecases.CreateTransfer(fromAcc, toAcc, 70.00)

	if fromAcc.Balance != 30.0 {
		t.Error("Expected", 30.0, "Got", fromAcc.Balance)
	}

	if toAcc.Balance != 90.0 {
		t.Error("Expected", 90.0, "Got", toAcc.Balance)
	}
}
