package usecases

import (
	"fmt"

	"banco_de_espana/entities"
)

// CreateAccount creates an account and persist in database
func CreateAccount(cl *entities.Client, b float64) (*entities.Account, error) {
	var err error

	if cl == nil {
		err = fmt.Errorf("Client should be passed")
	}

	return &entities.Account{0, cl, b}, err
}
