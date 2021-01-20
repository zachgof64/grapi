package main

import (
	"github.com/gorilla/mux"
	"github.com/zeuce/golang-api/routes"
	"net/http"
)

//ResponseStruct - Normal response 
type ResponseStruct struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}


//SetupDefaultHeaders - Adds some global headers
func SetupDefaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Accept", "application/json")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET POST PUT DELETE")
		next.ServeHTTP(w,r)
	})
}


//GetRouter - Returns mux router
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(SetupDefaultHeaders)
	r.HandleFunc("/status", routes.Status)
	if(GetConfig().Logging){
		SetupLogging(r)
	}
	return r
}
