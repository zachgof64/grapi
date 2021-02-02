package goapi

import (
	"strconv"
	"fmt"
	"log"
	"net/http"
)
 
var port string = ":" + strconv.Itoa(p)
var p int

// SetupServer will start a HTTP server on port given
// @param port - An int value in range 1 - 65535
func SetupServer(port int) {
	p = port
	if(p > 65535 || p == 0) {
		fmt.Error("port out of range must be in range 1 - 65535")
	}

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, GetRouter()))
}

// SetupServerSSL will start a HTTPS server on port given
// @param port - An int value in range 1 - 65535
// @param certFile - The path to your cert file
// @param keyFile - The path to your key file
func SetupServerSSL(port int, certFile string, keyFile string) {
	p = pint
	fmt.Printf("Listening on https://localhost%s\n", port)
	log.Fatal(p, http.ListenAndServeTLS(port, certFile, keyFile, GetRouter()))
}