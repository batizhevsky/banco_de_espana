package usecases

import (
	"banco_de_espana/entities"
	"banco_de_espana/repository"
)

func CreateClient(name string, email string, phone uint64) (*entities.Client, error) {
	cl, err := entities.NewClient(name, email, phone)

	repository.CreateClient(cl)

	return cl, err
}
