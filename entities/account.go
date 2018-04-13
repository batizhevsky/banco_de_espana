package entities

import "fmt"

// Account of the bank
type Account struct {
	ID      int
	Client  *Client
	Balance float64
}

// NewAccount ...
func NewAccount(cl *Client, b float64) (*Account, error) {
	var err error

	if cl == nil {
		err = fmt.Errorf("Client should be passed")
	}

	return &Account{0, cl, b}, err
}
