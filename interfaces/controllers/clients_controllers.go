package controllers

import (
	"banco_de_espana/entities"
	"banco_de_espana/presenters"
	"banco_de_espana/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ClientsIndex(w http.ResponseWriter, _ *http.Request) {
	list, err := repository.ListClients()
	if err != nil {
		renderError(err, w)
		return
	}

	renderJson(list, w)
}

func ClientShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 0, 64)
	if err != nil {
		renderError(err, w)
		return
	}

	client, err := repository.GetClient(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}

		renderError(err, w)
		return
	}

	renderJson(client, w)
}

func ClientCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		renderError(err, w)
		return
	}

	var payload presenters.ClientJSONPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		renderError(err, w)
		return
	}

	cl := payload.Data //.(map[string]interface{})
	builtCl, err := entities.NewClient(cl.Name, cl.Email, cl.Phone)

	if err != nil {
		renderBadRequest(err, w)
		return
	}

	// repository.CreateClient(builtCl)

	renderJson(builtCl, w)
}

func ClientUpdate(w http.ResponseWriter, r *http.Request) {

}

func ClientDelete(w http.ResponseWriter, r *http.Request) {

}

func renderError(err error, w http.ResponseWriter) {
	log.Fatal(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func renderBadRequest(err error, w http.ResponseWriter) {
	log.Fatal(err)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func renderJson(payload interface{}, w http.ResponseWriter) {
	json, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"data\": %s}", json)

	log.Printf("ClientsIndex serving response: %s", json)
}
