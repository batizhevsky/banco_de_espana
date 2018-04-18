package main

import (
	"fmt"

	"banco_de_espana/entities"
	"banco_de_espana/repository"
)

func main() {
	cl, err := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)

	if err != nil {
		panic("can't create a client")
	}

	repository.CreateClient(cl)

	fmt.Println(*cl)

	acc, err := entities.NewAccount(cl, 100.00)

	if err != nil {
		println(err)
		panic("can't create an account")
	}

	repository.CreateAccount(acc)

	fmt.Println(*acc)

	acc2, err := entities.NewAccount(cl, 50.00)

	if err != nil {
		println(err)
		panic("can't create an account")
	}

	repository.CreateAccount(acc2)

	fmt.Println(*acc2)
}
