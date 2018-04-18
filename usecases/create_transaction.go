package usecases

import "banco_de_espana/entities"

func CreateTransaction(acc *entities.Account, amount int64) error {
	var err error
	ChargeAccount(acc, amount)

	return err
}
