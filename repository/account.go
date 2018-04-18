package repository

import (
	"banco_de_espana/entities"
	"fmt"
)

// GetAccount ...
func GetAccount(id int64) *entities.Account {
	rows, err := connDB.Query("Select * FROM accounts where id = ? limit 1", id)

	checkErr(err)

	acc := entities.Account{}
	var clientID int64

	if rows.Next() {
		err = rows.Scan(&acc.ID, &clientID, &acc.Balance)
		checkErr(err)

		// acc.Client = GetClient(clientId) // Don't work
		return &acc
	}

	return nil
}

// CreateAccount ...
func CreateAccount(acc *entities.Account) {
	stmt, err := connDB.Prepare("INSERT INTO accounts(client_id, balance) values(?, ?)")
	checkErr(err)

	res, err := stmt.Exec(acc.Client.ID, acc.Balance)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	acc.ID = id
}

// UpdateBalance ...
func UpdateBalance(acc *entities.Account) {
	stmt, err := connDB.Prepare("update accounts set balance=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(acc.Balance, acc.ID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Printf("Balance was updated on account %v", affect)

}
