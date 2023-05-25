package main

import (
	"github.com/dchupakhin/go-study/http/gorilla"
	"log"
)

func main() {

	err := gorilla.StartServer(":8080")
	if err != nil {
		log.Fatal(err)
	}

	//err := simple.StartServer(":8080")
	//if err != nil {
	//	log.Fatal(err)
	//}
}
