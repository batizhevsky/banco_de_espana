package repository

import (
	"banco_de_espana/entities"
	"log"
)

// ClientDAL implements interfact for accessing to client's data
type ClientDAL interface {
	CreateClient(*entities.Client)
	GetClient(int64) (*entities.Client, error)
	ListClients() ([]entities.Client, error)
}

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
func GetClient(id int64) (*entities.Client, error) {
	row := ConnDB.QueryRow("Select * FROM clients where id = ? limit 1", id)

	cl := entities.Client{}
	err := row.Scan(&cl.ID, &cl.Name, &cl.Email, &cl.Phone)

	return &cl, err
}

// ListClients provides a list of all clients
func ListClients() ([]entities.Client, error) {
	rows, err := ConnDB.Query("Select * FROM clients")
	if err != nil {
		checkErr(err)
	}
	defer rows.Close()

	list := make([]entities.Client, 0)
	for rows.Next() {
		cl := entities.Client{}
		err = rows.Scan(&cl.ID, &cl.Name, &cl.Email, &cl.Phone)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, cl)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return list, err
}
