package repository

import (
	"banco_de_espana/entities"
	"testing"
)

func TestCreateClient(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	CreateClient(cl)

	if cl.ID == 0 {
		t.Error("Expected", "to be saved", "actually", "no")
	}

}

func TestGetClient(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	CreateClient(cl)

	dbCl := GetClient(cl.ID)

	if dbCl == nil || cl.Email != dbCl.Email || cl.Name != dbCl.Name || cl.Phone != dbCl.Phone {
		t.Error("Expected", "found", "actually", "no")
	}

	dbCl = GetClient(0)
	println(dbCl)

	if dbCl != nil {
		t.Error("Expected", "found nothing", "actually", dbCl)
	}

}
