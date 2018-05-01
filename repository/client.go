package repository

import (
	"banco_de_espana/entities"
	"errors"
	"fmt"
)

// CreateClient ...
func CreateClient(cl *entities.Client) {
	stmt, err := ConnDB.Prepare("INSERT INTO clients(name, email, phone) values(?, ?, ?)")

	checkErr(err)

	res, err := stmt.Exec(cl.Name, cl.Email, cl.Phone)

	checkErr(err)

	id, err := res.LastInsertId()

	checkErr(err)

	cl.ID = id
}

// GetClient ...
func GetClient(id int64) *entities.Client {
	rows, err := ConnDB.Query("Select * FROM clients where id = ? limit 1", id)

	checkErr(err)

	cl := entities.Client{}
	fmt.Println(rows)

	if rows.Next() {
		err = rows.Scan(&cl.ID, &cl.Name, &cl.Email, &cl.Phone)

		checkErr(err)
		return &cl
	}

	return nil
}

// Provide a list of all clients
func ListClients() ([]entities.Client, error) {
	return []entities.Client{}, errors.New("Not implemented")
}
