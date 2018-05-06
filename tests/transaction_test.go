package tests

import (
	"banco_de_espana/entities"
	"banco_de_espana/repository"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", "18005687625")
	repository.CreateClient(cl)

	acc, _ := entities.NewAccount(cl, 11000)
	repository.CreateAccount(acc)

	tr, _ := entities.NewTransaction(acc, 7000, "test purpose")
	repository.CreateTransaction(tr)

	if tr.ID == 0 || tr.Account.ID != acc.ID || tr.Amount != 7000 || tr.Subject != "test purpose" {
		t.Error("Expected", "to be saved", "actually", "no")
	}

}

func TestGetTransaction(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", "18005687625")
	repository.CreateClient(cl)

	acc, _ := entities.NewAccount(cl, 11000)
	repository.CreateAccount(acc)

	tr, _ := entities.NewTransaction(acc, 7000, "test purpose")
	repository.CreateTransaction(tr)

	dbTr := repository.GetTransaction(acc.ID)
	if dbTr == nil || dbTr.Amount != tr.Amount {
		t.Error("Expected", "found", "actually", "no")
	}

	dbTr = repository.GetTransaction(0)
	println(dbTr)

	if dbTr != nil {
		t.Error("Expected", "found nothing", "actually", dbTr)
	}
}
