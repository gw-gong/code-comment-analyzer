package middleware

import (
	"fmt"
	"log"
)

type extractorKey string

const (
	keyUserID extractorKey = "user_id"
)

type Extractor interface {
	setUserId(userID uint64)
	GetUserId() (uint64, error)
}

type extractedData map[extractorKey]interface{}

func newExtractedData() Extractor {
	return extractedData(make(map[extractorKey]interface{}))
}

func (e extractedData) setUserId(userID uint64) {
	if e == nil {
		log.Println("extractedData is not initialized")
		return
	}
	e[keyUserID] = userID
}

func (e extractedData) GetUserId() (uint64, error) {
	userID, ok := e[keyUserID]
	if !ok {
		return 0, fmt.Errorf("userId not found in extracted data")
	}
	return userID.(uint64), nil
}
