package usecases

import (
	"testing"

	"banco_de_espana/entities"
)

func TestChargeAccount(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	acc, _ := entities.NewAccount(cl, 100.00)

	ChargeAccount(acc, 90.00)

	if acc.Balance != 10.0 {
		t.Error("Expected", 10.0, "Got", acc.Balance)
	}
}

func TestCreditAccount(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	acc, _ := entities.NewAccount(cl, 10.0)

	CreditAccount(acc, 50.0)

	if acc.Balance != 60.0 {
		t.Error("Expected", 60.0, "Got", acc.Balance)
	}
}
