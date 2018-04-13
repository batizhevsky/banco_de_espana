package entities

import (
	"fmt"
	"time"
)

// Client of a bank
type Client struct {
	ID        int64
	Name      string
	Email     string
	Phone     uint64
	CreatedAt time.Time
}

// NewClient initialize a new client
func NewClient(name string, email string, phone uint64) (*Client, error) {
	var err error

	if name == "" {
		err = fmt.Errorf("Name should have a value")
	}

	if email == "" {
		err = fmt.Errorf("E-Mail should have a value")
	}

	if phone == 0 {
		err = fmt.Errorf("Phone should have a value")
	}

	return &Client{0, name, email, phone, time.Time{}}, err
}
