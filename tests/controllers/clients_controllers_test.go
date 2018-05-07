package controllers_test

import (
	"banco_de_espana/entities"
	"banco_de_espana/interfaces/controllers"
	"banco_de_espana/repository"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestClientsIndex(t *testing.T) {
	// tests.ClearTable()

	cl1, _ := entities.NewClient("John Jonson", "john@wallstreet.com", "+18005687625")
	cl2, _ := entities.NewClient("Jack Jackson", "jack@wallstreet.com", "+18005687626")
	repository.CreateClient(cl1)
	repository.CreateClient(cl2)

	request := httptest.NewRequest("GET", "/clients", nil)
	recorder := httptest.NewRecorder()

	controllers.ClientsIndex(recorder, request)

	resp := recorder.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code ok, got %v\n", resp.StatusCode)
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("Expected content type is json, got %v\n", ct)
	}

	if string(body) != `{data: [{"id":1,"Name":"test","Email":"test@go.com","Phone":"299932323"}]}` {
		t.Errorf("Unexpected body. Got %s\n", body)
	}

}
