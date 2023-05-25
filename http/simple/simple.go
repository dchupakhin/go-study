package simple

import (
	"github.com/dchupakhin/go-study/http/simple/server/handler"
	"net/http"
)

const MARK = "Simple"

func StartServer(address string) error {
	http.HandleFunc("/", handler.IndexHandler(MARK))
	http.HandleFunc("/home", handler.HomeHandler(MARK))
	http.HandleFunc("/hello/", handler.HelloHandler(MARK))

	return http.ListenAndServe(address, nil)
}
