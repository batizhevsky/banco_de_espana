package usecases

import "banco_de_espana/entities"

func CreateTransfer(from *entities.Account, to *entities.Account, amount int64) error {
	var err error
	ChargeAccount(from, amount)
	CreditAccount(to, amount)

	return err
}
