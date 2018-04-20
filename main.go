package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexHandler struct {
	routes []*route
}

func (h *RegexHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

func (h *RegexHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func main() {
	http.HandleFunc("/test", sayhellotest)
	http.HandleFunc("/", sayhelloName)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhellotest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Other!!!!!")
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println(r.Form["url_long"])
	fmt.Fprintf(w, "Hello astaxie!")
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	counter++
	println(counter)
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println(r.Form["url_long"])
}
