package usecases

import (
	"banco_de_espana/entities"
	"banco_de_espana/repository"
)

// ChargeAccount ...
func ChargeAccount(acc *entities.Account, amount int64) error {
	var err error

	acc.Balance -= amount

	createTransaction(acc, -amount)
	repository.UpdateBalance(acc)

	return err
}

// CreditAccount ...
func CreditAccount(acc *entities.Account, amount int64) error {
	var err error

	acc.Balance += amount

	createTransaction(acc, amount)
	repository.UpdateBalance(acc)

	return err
}

func createTransaction(acc *entities.Account, amount int64) {
	tr, err := entities.NewTransaction(acc, amount, "")

	if err == nil {
		repository.CreateTransaction(tr)
	} else {
		panic(err)
	}

}
