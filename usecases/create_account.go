package usecases

import (
	"banco_de_espana/repository"

	"banco_de_espana/entities"
)

// CreateAccount creates an account and persist in database
func CreateAccount(cl *entities.Client, b int64) (*entities.Account, error) {
	// var err error

	// if cl == nil {
	// 	err = fmt.Errorf("Client should be passed")
	// 	return nil, err
	// }

	acc, err := entities.NewAccount(cl, b)

	repository.CreateAccount(acc)

	return acc, err
}
