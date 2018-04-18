package repository

import (
	"banco_de_espana/entities"
)

// GetAccount ...
func GetAccount(id int64) *entities.Account {
	rows, err := connDB.Query("Select * FROM accounts where id = ? limit 1", id)

	checkErr(err)

	acc := entities.Account{}
	var clientId int64

	if rows.Next() {
		err = rows.Scan(&acc.ID, &clientId, &acc.Balance)

		checkErr(err)

		// acc.Client = GetClient(clientId) // Don't work
		return &acc
	}

	return nil
}

func CreateAccount(acc *entities.Account) {
	stmt, err := connDB.Prepare("INSERT INTO accounts(client_id, balance) values(?, ?)")

	checkErr(err)

	res, err := stmt.Exec(acc.Client.ID, acc.Balance)

	checkErr(err)

	id, err := res.LastInsertId()

	checkErr(err)

	acc.ID = id

}
