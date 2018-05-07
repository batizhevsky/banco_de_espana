package repository

import (
	"banco_de_espana/entities"
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

	if rows.Next() {
		err = rows.Scan(&cl.ID, &cl.Name, &cl.Email, &cl.Phone)

		checkErr(err)
		return &cl
	}

	return nil
}

// ListClients provides a list of all clients
func ListClients() ([]entities.Client, error) {
	rows, err := ConnDB.Query("Select * FROM clients")

	list := make([]entities.Client, 0)
	if rows.Next() {
		cl := entities.Client{}
		err = rows.Scan(&cl.ID, &cl.Name, &cl.Email, &cl.Phone)
		list = append(list, cl)
	}
	return list, err
}
