package repository

import (
	"banco_de_espana/entities"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateClient(cl *entities.Client) {
	stmt, err := connDB.Prepare("INSERT INTO clients(name, email, phone) values(?, ?, ?)")

	checkErr(err)

	res, err := stmt.Exec(cl.Name, cl.Email, cl.Phone)

	checkErr(err)

	id, err := res.LastInsertId()

	checkErr(err)

	cl.ID = id
}

func GetClient(id int64) *entities.Client {
	rows, err := connDB.Query("Select * FROM clients where id = ? limit 1", id)

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
