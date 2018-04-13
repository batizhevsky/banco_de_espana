package usecases

import "banco_de_espana/entities"

func CreateTransaction(from *entities.Account, to *entities.Account, amount float64) error {
	var err error
	ChargeAccount(from, amount)
	CreditAccount(to, amount)

	return err
}
