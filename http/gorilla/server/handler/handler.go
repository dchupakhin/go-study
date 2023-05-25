package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dchupakhin/go-study/model"
	"github.com/gorilla/mux"
)

type NotFoundHandler struct {
}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 >>> %s", r.URL.Path)
	http.NotFound(w, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Path: %s", r.URL.Path)
	_, err := w.Write([]byte("Home >>> Gorilla"))
	if err != nil {
		log.Fatal(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Path: %s", r.URL.Path)
	_, err := w.Write([]byte("Index >>> Gorilla"))
	if err != nil {
		log.Fatal(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	log.Printf("Path: %s", r.URL.Path)
	_, err := w.Write([]byte(fmt.Sprintf("Hello, %s! >>> Gorilla", name)))
	if err != nil {
		log.Fatal(err)
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["id"]
	log.Printf("Path: %s", r.URL.Path)
	_, err := w.Write([]byte(fmt.Sprintf("Product id: %s >>> Gorilla", name)))
	if err != nil {
		log.Fatal(err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Path: %s", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	responseUser := model.User{ID: 1, Name: "Godev", Email: "sdf@sdf.ru", Password: "12345"}
	encodeErr := json.NewEncoder(w).Encode(responseUser)
	if encodeErr != nil {
		panic(encodeErr)
	}
	var requestUser model.User
	decodeErr := json.NewDecoder(r.Body).Decode(&requestUser)
	if decodeErr != nil {
		panic(decodeErr)
	}
	log.Printf("%#v", requestUser)
}
