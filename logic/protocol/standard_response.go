package protocol

import (
	"encoding/json"
	"log"
	"net/http"
)

type option func(*optionParams)

type optionParams struct {
	language string
	data     interface{}
}

func WithLanguage(language string) option {
	return func(params *optionParams) {
		params.language = language
	}
}

func WithData(data interface{}) option {
	return func(params *optionParams) {
		params.data = data
	}
}

// ps: 由于项目是重构的，所以这里会多一个language字段，不是很合理，但是现在前端不修改的情况下，只能这么加进去了。
// 实际上只需要关心 status, msg, data 即可

type responseFormat struct {
	Status   int         `json:"status"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
	Language string      `json:"language,omitempty"`
}

func httpResponse(w http.ResponseWriter, httpStatusCode int, status int, msg string, opts ...option) {
	optParams := &optionParams{}
	for _, opt := range opts {
		opt(optParams)
	}

	response := responseFormat{
		Status:   status,
		Msg:      msg,
		Data:     optParams.data,
		Language: optParams.language,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("json encoding error: %v", err)
		return
	}
}

func HttpResponseSuccess(w http.ResponseWriter, httpStatusCode int, msg string, opts ...option) {
	httpResponse(w, httpStatusCode, StatusSuccess, msg, opts...)
}

func HttpResponseFail(w http.ResponseWriter, httpStatusCode int, errorCode int, msg string) {
	httpResponse(w, httpStatusCode, errorCode, msg)
}
