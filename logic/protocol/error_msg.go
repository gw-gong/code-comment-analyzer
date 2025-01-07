package protocol

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ErrorMsgInvalidToken string = "invalid token"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func HandleError(w http.ResponseWriter, errorCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	errorResponse := ErrorResponse{
		Status: "error",
		Code:   errorCode,
		Msg:    fmt.Sprintf("%v", err),
	}
	json.NewEncoder(w).Encode(errorResponse)
}
