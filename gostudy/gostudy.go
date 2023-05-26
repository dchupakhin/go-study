package main

import (
	"github.com/dchupakhin/go-study/http/server"
)

func main() {

	if err := server.StartServer(":8080"); err != nil {
		panic(err)
	}
}
