package goapi

import (
	"github.com/gorilla/mux"
)

var router *mux.Router

//ResponseStruct - Normal response with no data
type ResponseStruct struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

//SetupRouter - Setup Mux router
func SetupRouter(defaultHeaders ...string) {
	router = mux.NewRouter()
}

//GetRouter - Returns mux router
func GetRouter() *mux.Router {
	return router
}