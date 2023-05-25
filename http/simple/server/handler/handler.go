package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func HomeHandler(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/home" {
			http.NotFound(w, r)
			log.Printf("404 >>> path: %s mark: %s", r.URL.Path, msg)
			return
		}
		log.Printf("Path: %s", r.URL.Path)
		_, err := w.Write([]byte(fmt.Sprintf("Home >>> %s", msg)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func IndexHandler(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != `/` {
			http.NotFound(w, r)
			log.Printf("404 >>> path: %s mark: %s", r.URL.Path, msg)
			return
		}
		log.Printf("Path: %s", r.URL.Path)
		_, err := w.Write([]byte(fmt.Sprintf("Index >>> %s", msg)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func HelloHandler(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if pathIsNotValid(r.URL.Path, `/hello/\w+$`) {
			http.NotFound(w, r)
			log.Printf("404 >>> path: %s mark: %s", r.URL.Path, msg)
			return
		}
		name := strings.Split(r.URL.Path, "/")[2]
		log.Printf("Path: %s", r.URL.Path)
		_, err := w.Write([]byte(fmt.Sprintf("Hello, %s! >>> %s", name, msg)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func pathIsNotValid(path string, regexPattern string) bool {
	pathRegex := regexp.MustCompile(regexPattern)
	return !pathRegex.Match([]byte(path))
}
