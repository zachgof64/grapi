package goapi

import (
	"github.com/gorilla/mux"
)

// ResponseStruct is a struct for to help prevent code duplication
type ResponseStruct struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}
