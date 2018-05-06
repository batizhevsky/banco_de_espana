package infrastructure

import (
	"log"
	"net/http"
	"os"
)

func RunWebserver() {
	r := BindRoutes()

	src := &http.Server{
		Addr:    "0.0.0.0:9090",
		Handler: logRequest(r),
	}
	if err := src.ListenAndServe(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
