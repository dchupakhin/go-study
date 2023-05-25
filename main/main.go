package main

import "net/http"

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home", homeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := w.Write([]byte("Hello, World"))
	if err != nil {
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" && r.URL.Path != "/home/" {
		http.NotFound(w, r)
		return
	}
	_, err := w.Write([]byte("Hello, Home"))
	if err != nil {
		return
	}
}
