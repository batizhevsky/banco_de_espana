package tests

import (
	"banco_de_espana/entities"
	"banco_de_espana/repository"
	"testing"
)

func TestCreateClient(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	repository.CreateClient(cl)

	if cl.ID == 0 {
		t.Error("Expected", "to be saved", "actually", "no")
	}

}

func TestGetClient(t *testing.T) {
	cl, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	repository.CreateClient(cl)

	dbCl := repository.GetClient(cl.ID)

	if dbCl == nil || cl.Email != dbCl.Email || cl.Name != dbCl.Name || cl.Phone != dbCl.Phone {
		t.Error("Expected", "found", "actually", "no")
	}

	dbCl = repository.GetClient(0)

	if dbCl != nil {
		t.Error("Expected", "found nothing", "actually", dbCl)
	}

}

func TestListClients(t *testing.T) {
	cl1, _ := entities.NewClient("John Jonson", "john@wallstreet.com", 18005687625)
	cl2, _ := entities.NewClient("Jack Jackson", "jack@wallstreet.com", 18005687626)
	repository.CreateClient(cl1)
	repository.CreateClient(cl2)

	list, err := repository.ListClients()

	if err != nil {
		t.Errorf("Client list should be fetched without errors, but got %v", err)
	}

	if len(list) != 2 {
		t.Error("Expected", "2 items", "got", list)
	}

}
