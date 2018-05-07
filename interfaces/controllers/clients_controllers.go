package controllers

import (
	"banco_de_espana/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ClientsIndex(w http.ResponseWriter, _ *http.Request) {
	list, err := repository.ListClients()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something went wrong: %v\n", err)
		return
	}

	json, _ := json.Marshal(list)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{data: %s}", json)

	log.Printf("ClientsIndex serving response: %s", json)
	w.WriteHeader(http.StatusOK)
}

func ClientShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 0, 64)

	if err == nil && id != 0 {
		w.WriteHeader(http.StatusOK)

		client := repository.GetClient(id)
		json, _ := json.Marshal(client)

		fmt.Fprintf(w, "{data: %s}", json)
	}
}

func ClientCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	println(vars)
}

func ClientUpdate(w http.ResponseWriter, r *http.Request) {

}

func ClientDelete(w http.ResponseWriter, r *http.Request) {

}
