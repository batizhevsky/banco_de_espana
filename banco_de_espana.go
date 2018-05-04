package main

import (
	"banco_de_espana/infrastructure"
	"banco_de_espana/usecases"
)

func main() {
	usecases.CreateClient("test", "test@go.com", 299932323)

	infrastructure.RunWebserver()
}
