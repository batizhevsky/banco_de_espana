package repository

import (
	"time"

	"banco_de_espana/entities"

	_ "github.com/mattn/go-sqlite3"
)

func CreateClient(cl *entities.Client) {
	stmt, err := DBConnection().Prepare("INSERT INTO clients(name, email, phone, created_at) values(?, ?, ?, ?)")

	checkErr(err)

	res, err := stmt.Exec(cl.Name, cl.Email, cl.Phone, time.Now())

	checkErr(err)

	id, err := res.LastInsertId()

	checkErr(err)

	cl.ID = id
}

func GetClient(id int64) *entities.Client {
	rows, err := DBConnection().Query("Select * FROM clients")
	cl := entities.Client{}

	err = rows.Scan(&cl.ID, &cl.Name, &cl.Email, &cl.Phone, &cl.CreatedAt)

	return cl
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
