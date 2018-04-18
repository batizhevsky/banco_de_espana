package usecases

import (
	"banco_de_espana/entities"
)

// ChargeAccount ...
func ChargeAccount(acc *entities.Account, amount int64) error {
	var err error

	acc.Balance -= amount

	return err
}

// CreditAccount ...
func CreditAccount(acc *entities.Account, amount int64) error {
	var err error

	acc.Balance += amount

	return err
}
