package gorilla

import (
	"net/http"

	"github.com/dchupakhin/go-study/http/gorilla/server/handler"

	"github.com/gorilla/mux"
)

func StartServer(address string) error {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.IndexHandler)
	r.HandleFunc("/home", handler.HomeHandler)
	r.HandleFunc("/hello/{name}", handler.HelloHandler)
	r.HandleFunc(`/product/{id:\d+}`, handler.ProductHandler)
	r.HandleFunc("/form", handler.FormHandler).Methods("POST", "PUT")
	r.NotFoundHandler = handler.NotFoundHandler{}

	return http.ListenAndServe(address, r)
}
