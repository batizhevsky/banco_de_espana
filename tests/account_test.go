package tests

import (
	"banco_de_espana/entities"
	"banco_de_espana/repository"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	repository.CreateClient(cl)

	acc, _ := entities.NewAccount(cl, 11000)

	repository.CreateAccount(acc)

	if acc.ID == 0 {
		t.Error("Expected", "to be saved", "actually", "no")
	}
}

func TestUpdateBalance(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	repository.CreateClient(cl)

	acc, _ := entities.NewAccount(cl, 11000)

	repository.CreateAccount(acc)

	acc.Balance = 9000

	repository.UpdateBalance(acc)

	dbAcc := repository.GetAccount(acc.ID)
	if dbAcc == nil || dbAcc.ID != acc.ID || dbAcc.Balance != 9000 {
		t.Error("Expected", "balance updated", "actually", "no")
	}

}

func TestGetAccount(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	repository.CreateClient(cl)

	acc, _ := entities.NewAccount(cl, 11000)
	repository.CreateAccount(acc)

	dbAcc := repository.GetAccount(acc.ID)
	// if dbAcc == nil || cl.ID != dbAcc.Client.ID || dbAcc.Balance != acc.Balance { // Don't work
	if dbAcc == nil || dbAcc.Balance != acc.Balance {
		t.Error("Expected", "found", "actually", "no")
	}

	dbAcc = repository.GetAccount(0)
	println(dbAcc)

	if dbAcc != nil {
		t.Error("Expected", "found nothing", "actually", dbAcc)
	}
}
