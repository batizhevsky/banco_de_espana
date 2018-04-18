package usecases

import (
	"testing"
)

func TestCreateAccount(t *testing.T) {
	cl, err := CreateClient("tester", "email@do.com", 18005687625)

	if err != nil {
		t.Error("Expected", "record should be saved", "but got", err)
	}
}
