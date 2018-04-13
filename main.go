package main

import (
	"fmt"

	"banco_de_espana/entities"
)

func main() {
	cl, err := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)

	if err != nil {
		panic("can't create a client")
	}

	fmt.Println(*cl)
	acc, err := entities.NewAccount(cl, 100.00)

	if err != nil {
		println(err)
		panic("can't create an account")
	}

	fmt.Println(*acc)
}
