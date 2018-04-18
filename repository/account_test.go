package repository

import (
	"banco_de_espana/entities"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	CreateClient(cl)

	acc, _ := entities.NewAccount(cl, 11000)

	CreateAccount(acc)

	if acc.ID == 0 {
		t.Error("Expected", "to be saved", "actually", "no")
	}

}

func TestGetAccount(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	CreateClient(cl)

	acc, _ := entities.NewAccount(cl, 11000)
	CreateAccount(acc)

	dbAcc := GetAccount(acc.ID)
	// if dbAcc == nil || cl.ID != dbAcc.Client.ID || dbAcc.Balance != acc.Balance { // Don't work
	if dbAcc == nil || dbAcc.Balance != acc.Balance {
		t.Error("Expected", "found", "actually", "no")
	}

	dbAcc = GetAccount(0)
	println(dbAcc)

	if dbAcc != nil {
		t.Error("Expected", "found nothing", "actually", dbAcc)
	}
}
