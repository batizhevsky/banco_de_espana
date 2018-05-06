package usecases

import (
	"banco_de_espana/client/infrastucture"
	"banco_de_espana/entities"
	"banco_de_espana/presenters"
)

func CreateClient(name, email, phone string) error {
	http := infrastucture.NewHTTPClient()

	client, err := entities.NewClient(name, email, phone)
	if err != nil {
		return err
	}

	json, err := presenters.SerializeJSON(client)
	if err != nil {
		return err
	}

	http.Post("/clients", json)

	return nil
}
