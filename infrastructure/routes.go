package infrastructure

import (
	"banco_de_espana/interfaces/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// BindRoutes returns binded routes
func BindRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/clients", controllers.ClientsIndex).Methods("GET")
	r.HandleFunc("/clients", controllers.ClientCreate).Methods("POST")
	r.HandleFunc("/clients/{id}", controllers.ClientShow).Methods("GET")
	r.HandleFunc("/clients/{id}", controllers.ClientUpdate).Methods("PUT")
	r.HandleFunc("/clients/{id}", controllers.ClientDelete).Methods("DELETE")

	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Customer!")
}
