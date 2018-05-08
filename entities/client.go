package entities

import (
	"fmt"
)

// Client of a bank
type Client struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// NewClient initialize a new client
func NewClient(name string, email string, phone string) (*Client, error) {
	var err error

	if name == "" {
		err = fmt.Errorf("Name should have a value")
	}

	if email == "" {
		err = fmt.Errorf("E-Mail should have a value")
	}

	if phone == "" {
		err = fmt.Errorf("Phone should have a value")
	}

	return &Client{0, name, email, phone}, err
}
