package protocol

import (
	"encoding/json"
	"log"
	"net/http"
)

// ps: 由于项目是重构的，所以这里会多一个language字段，不是很合理，但是现在前端不修改的情况下，只能这么加进去了。
// 这边使用了"..."参数形式，来降低项目中的影响，直接忽略即可，只有少数接口需要用到。
// 实际上只需要关心 status, msg, data 即可

type responseFormat struct {
	Status   int         `json:"status"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
	Language string      `json:"language,omitempty"`
}

func httpResponse(w http.ResponseWriter, httpStatusCode int, status int, msg string, data interface{}, language ...string) {
	validLanguage := ""
	if len(language) == 1 {
		validLanguage = language[0]
	}
	response := responseFormat{
		Status:   status,
		Msg:      msg,
		Data:     data,
		Language: validLanguage,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("json encoding error: %v", err)
		return
	}
}

func HttpResponseSuccess(w http.ResponseWriter, httpStatusCode int, msg string, data interface{}, language ...string) {
	httpResponse(w, httpStatusCode, StatusSuccess, msg, data, language...)
}

func HttpResponseFail(w http.ResponseWriter, httpStatusCode int, errorCode int, msg string) {
	httpResponse(w, httpStatusCode, errorCode, msg, nil)
}
