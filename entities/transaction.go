package entities

import (
	"fmt"
	"time"
)

// Transaction ...
type Transaction struct {
	ID       int64
	Account  *Account
	Amount   int64
	Subject  string
	DateTime time.Time
}

// NewTransaction initilizes transaction
func NewTransaction(acc *Account, amount int64, subj string) (*Transaction, error) {
	var err error

	if acc == nil {
		err = fmt.Errorf("Account should be setted")
	}

	if amount == 0 {
		err = fmt.Errorf("to account should be setted")
	}

	return &Transaction{0, acc, amount, subj, time.Now()}, err
}
