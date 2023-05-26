package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(address string) error {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/home", HomeHandler)
	r.HandleFunc("/hello/{name}", HelloHandler)
	r.HandleFunc(`/product/{id:\d+}`, ProductHandler)
	r.HandleFunc("/form", FormHandler).Methods("POST", "PUT")
	r.NotFoundHandler = NotFoundHandler{}

	return http.ListenAndServe(address, r)
}
