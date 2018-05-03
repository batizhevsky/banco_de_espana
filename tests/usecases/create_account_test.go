package usecases_test

import (
	"banco_de_espana/usecases"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	cl, err := usecases.CreateClient("tester", "email@do.com", 18005687625)

	if err != nil {
		t.Error("Expected", "record should be saved", "but got", err)
	}

	if cl.ID == 0 {
		t.Error("Expected", "has an ID", "but got", 0)
	}
}
