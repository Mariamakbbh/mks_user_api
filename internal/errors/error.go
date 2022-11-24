package errors

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ErrorMsg struct {
	Body Error `json:"body"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetGenericError(code int, message string) ErrorMsg {
	err := Error{Code: code, Message: message}
	return ErrorMsg{Body: err}
}

func ProcessError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	errRes := GetGenericError(statusCode, err.Error())
	log.Error(err)
	json.NewEncoder(w).Encode(errRes.Body)
}
