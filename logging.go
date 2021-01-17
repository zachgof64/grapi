package main

import (
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"strings"
	"fmt"
)

//SetupLogging ...
func SetupLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile(GetConfig().LogFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		Check(err)
		logger := log.New(f, GetConfig().LogPrefix + " ", log.LstdFlags)
		logger.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		fmt.Printf("%s %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}


//CreateLogFile ...
func CreateLogFile() {
	lf := strings.Split(GetConfig().LogFile, "/")
	if _, err := os.Stat(GetConfig().LogFile); os.IsNotExist(err) {
		os.Mkdir(lf[0], 0755)
		ioutil.WriteFile(GetConfig().LogFile, []byte(""), 0744)	
	}
}