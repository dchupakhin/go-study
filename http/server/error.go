package server

import (
	"encoding/json"

	. "github.com/dchupakhin/go-study/logger"
)

type HandlerError struct {
	Handler string `json:"handler"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (he HandlerError) Error() string {
	value, err := json.Marshal(he)
	if err != nil {
		ErrorLog.Println(err)
	}

	return string(value)
}
