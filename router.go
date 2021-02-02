package goapi

import (
	"github.com/gorilla/mux"
)

//ResponseStruct - Normal response with no data
type ResponseStruct struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}
