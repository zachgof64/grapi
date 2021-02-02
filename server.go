package goapi

import (
	"strconv"
	"fmt"
	"log"
	"net/http"
)
 
var port string = ":" + strconv.Itoa(p)
var p int

//SetupServer - Start HTTP server with mux
func SetupServer(port int) {
	p = port
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, GetRouter()))
}

//SetupServerSSL - Start HTTPS server with mux
func SetupServerSSL(pint int, certFile string, keyFile string) {
	p = pint
	fmt.Printf("Listening on https://localhost%s\n", port)
	log.Fatal(p, http.ListenAndServeTLS(port, certFile, keyFile, GetRouter()))
}