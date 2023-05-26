package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/dchupakhin/go-study/logger"
	"github.com/gorilla/mux"
)

type NotFoundHandler struct {
}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ErrorLog.Printf("404 >>> %s %s", r.Method, r.URL.Path)
	http.NotFound(w, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	InfoLog.Printf("request >>> %s %s", r.Method, r.URL.Path)
	if _, err := w.Write([]byte("Home page")); err != nil {
		ErrorLog.Println(HandlerError{"HomeHandler", http.StatusInternalServerError, err.Error()})
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	InfoLog.Printf("request >>> %s %s", r.Method, r.URL.Path)
	if _, err := w.Write([]byte("Index page")); err != nil {
		ErrorLog.Println(HandlerError{"IndexHandler", http.StatusInternalServerError, err.Error()})
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	InfoLog.Printf("request >>> %s %s", r.Method, r.URL.Path)
	name := mux.Vars(r)["name"]
	if _, err := w.Write([]byte(fmt.Sprintf("Hello, %s!", name))); err != nil {
		ErrorLog.Println(HandlerError{"HelloHandler", http.StatusInternalServerError, err.Error()})
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	InfoLog.Printf("request >>> %s %s", r.Method, r.URL.Path)
	name := mux.Vars(r)["id"]
	if _, err := w.Write([]byte(fmt.Sprintf("Product id: %s", name))); err != nil {
		ErrorLog.Println(HandlerError{"ProductHandler", http.StatusInternalServerError, err.Error()})
		return
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	InfoLog.Printf("request >>> %s %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	var requestUser User
	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		ErrorLog.Println(HandlerError{"FormHandler", http.StatusBadRequest, err.Error()})
		return
	}
	responseUser := User{ID: 1, Name: "Godev", Email: "sdf@sdf.ru", Password: "12345"}
	if err := json.NewEncoder(w).Encode(responseUser); err != nil {
		ErrorLog.Println(HandlerError{"FormHandler", http.StatusInternalServerError, err.Error()})
		return
	}
}
