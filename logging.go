package goapi

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
	"github.com/zeuce/golang-utils"	
	"fmt"
	"io/ioutil"
)

var fullLogFile string
var logP string

//SetupLogging - Setups logging
func SetupLogging(r *mux.Router, logDir string, logFile string, logPrefix string) {
	r.Use(LogHandler)
	fullLogFile = logDir + "/" + logFile
	logP = logPrefix
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
		
	}
	if _, err := os.Stat(fullLogFile); os.IsNotExist(err) {
		ioutil.WriteFile(fullLogFile, []byte(""), 0744)
	}
}

//LogHandler - Adds log to file and outputs it to console
func LogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile(fullLogFile, os.O_APPEND | os.O_RDWR, 0744)
		defer f.Close()
		utils.Check(err)
		logger := log.New(f, logP + " ", log.LstdFlags)
		logger.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		fmt.Printf("%s %s\n" ,  r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
