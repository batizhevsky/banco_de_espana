package entities

import "fmt"

// Account of the bank
type Account struct {
	ID      int64
	Client  *Client
	Balance int64
}

// NewAccount ...
func NewAccount(cl *Client, b int64) (*Account, error) {
	var err error

	if cl == nil {
		err = fmt.Errorf("Client should be passed")
	}

	return &Account{0, cl, b}, err
}
