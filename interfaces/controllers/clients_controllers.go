package controllers

import (
	"banco_de_espana/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ClientsIndex(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

}

func ClientShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 0, 64)

	if err == nil && id != 0 {
		w.WriteHeader(http.StatusOK)

		client := repository.GetClient(id)
		json, _ := json.Marshal(client)

		fmt.Fprintf(w, "{data: [%s]}", json)
	}
}

func ClientCreate(w http.ResponseWriter, r *http.Request) {

}

func ClientUpdate(w http.ResponseWriter, r *http.Request) {

}

func ClientDelete(w http.ResponseWriter, r *http.Request) {

}
