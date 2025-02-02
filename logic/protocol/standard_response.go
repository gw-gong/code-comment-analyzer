package protocol

import (
	"encoding/json"
	"log"
	"net/http"
)

type responseFormat struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func httpResponse(w http.ResponseWriter, httpStatusCode int, status int, msg string, data interface{}) {
	response := responseFormat{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("json encoding error: %v", err)
		return
	}
}

func HttpResponseSuccess(w http.ResponseWriter, httpStatusCode int, msg string, data interface{}) {
	httpResponse(w, httpStatusCode, 0, msg, data)
}

func HttpResponseFail(w http.ResponseWriter, httpStatusCode int, msg string) {
	httpResponse(w, httpStatusCode, 1, msg, nil)
}
