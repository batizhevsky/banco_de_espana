package usecases

import (
	"banco_de_espana/entities"
	"log"
)

func SendWelcomeEmailAsync(c *entities.Client) {
	go SendWelcomeEmail(c)
}

func SendWelcomeEmail(c *entities.Client) error {
	log.Printf("Sending email to client %v", c)
	return nil
}
