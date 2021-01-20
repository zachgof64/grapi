package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
	"github.com/zeuce/golang-api/utils"	
	"fmt"
	"io/ioutil"
)

var fullLogFile string = GetConfig().LogDir + "/" + GetConfig().LogFile

//LogHandler - Adds log to file and outputs it to console
func LogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile(fullLogFile, os.O_APPEND | os.O_RDWR, 0744)
		defer f.Close()
		utils.Check(err)
		logger := log.New(f, GetConfig().LogPrefix + " ", log.LstdFlags)
		logger.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		fmt.Printf("%s %s\n" ,  r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

//SetupLogging - Setups logging
func SetupLogging(r *mux.Router) {
	r.Use(LogHandler)
	if _, err := os.Stat(GetConfig().LogDir); os.IsNotExist(err) {
		os.MkdirAll(GetConfig().LogDir, 0755)
		
	}
	if _, err := os.Stat(fullLogFile); os.IsNotExist(err) {
		ioutil.WriteFile(fullLogFile, []byte(""), 0744)
	}
}