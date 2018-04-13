package repository

import "banco_de_espana/entities"

const (
	dbName = "accounts"
)

func GetAccount(id int64) *entities.Account {
	rows, err := DBConnection().Query("Select * FROM " + dbName)
	acc = entities.Account{}

	return acc
}
