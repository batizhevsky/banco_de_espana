package usecases

import (
	"banco_de_espana/entities"
	"banco_de_espana/repository"
)

func CreateClient(name string, email string, phone string) (*entities.Client, error) {
	cl, err := entities.NewClient(name, email, phone)

	if err != nil {
		return cl, err
	}

	repository.CreateClient(cl)
	SendWelcomeEmailAsync(cl)

	return cl, nil
}
