package main

import (
	"github.com/gorilla/mux"
	"github.com/zeuce/golang-api/routes"
	"net/http"
)

//SetupDefaultHeaders ...
func SetupDefaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Accept", "application/json")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET POST PUT DELETE")
		next.ServeHTTP(w,r)
	})
}

//GetRouter ...
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(SetupDefaultHeaders)
	if(GetConfig().Logging){
		CreateLogFile()
		r.Use(SetupLogging)
	}
	r.HandleFunc("/status", routes.Status)
	return r
}