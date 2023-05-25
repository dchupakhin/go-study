package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type NotFoundHandler struct {
}

func HomeHandler(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path: %s", r.URL.Path)
		_, err := w.Write([]byte(fmt.Sprintf("Home >>> %s", msg)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func IndexHandler(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path: %s", r.URL.Path)
		_, err := w.Write([]byte(fmt.Sprintf("Index >>> %s", msg)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func HelloHandler(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		log.Printf("Path: %s", r.URL.Path)
		_, err := w.Write([]byte(fmt.Sprintf("Hello, %s! >>> %s", name, msg)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ProductHandler(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["id"]
		log.Printf("Path: %s", r.URL.Path)
		_, err := w.Write([]byte(fmt.Sprintf("Product id: %s >>> %s", name, msg)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func FormHandler(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path: %s", r.URL.Path)
		_, err := w.Write([]byte(fmt.Sprintf("Form >>> %s", msg)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 >>> %s", r.URL.Path)
	http.NotFound(w, r)
}
