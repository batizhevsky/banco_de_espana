package main

import (
	"banco_de_espana/repository"
	"banco_de_espana/usecases"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	usecases.CreateClient("test", "test@go.com", 299932323)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/clients", ClientsHandler).Methods("GET")
	r.HandleFunc("/clients/{id}", ClientHandler).Methods("GET")

	http.Handle("/", r)

	src := &http.Server{
		Addr:    "0.0.0.0:9090",
		Handler: logRequest(r),
	}
	src.ListenAndServe()

	// func() {
	// 	if err := src.ListenAndServe(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Customer!")
}

func ClientsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{data: []}")
}

func ClientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 0, 64)

	if err == nil && id != 0 {
		w.WriteHeader(http.StatusOK)

		client := repository.GetClient(id)
		json, _ := json.Marshal(client)

		fmt.Fprintf(w, "{data: [%s]}", json)
	}
}
