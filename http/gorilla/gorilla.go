package gorilla

import (
	"github.com/dchupakhin/go-study/http/gorilla/server/handler"
	"net/http"

	"github.com/gorilla/mux"
)

const MARK = "Gorilla"

func StartServer(address string) error {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.IndexHandler(MARK))
	r.HandleFunc("/home", handler.HomeHandler(MARK))
	r.HandleFunc("/hello/{name}", handler.HelloHandler(MARK))
	r.HandleFunc(`/product/{id:\d+}`, handler.ProductHandler(MARK))
	r.HandleFunc("/form", handler.FormHandler(MARK)).Methods("POST", "PUT")
	r.NotFoundHandler = handler.NotFoundHandler{}

	return http.ListenAndServe(address, r)
}
