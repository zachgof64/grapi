package grapi

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Router for internal use only
var Router *mux.Router = mux.NewRouter()

// Serve will start a HTTP server on port given
// @param router - Your mux router
// @param port - An int value in range 1 - 65535
func Serve(port int) {
	p := ":" + strconv.Itoa(port)
	if port > 65535 || port < 1 {
		fmt.Println(errors.New("port out of range must be in range 1 - 65535"))
	}
	Router.Use(defaultHeaderHandler)
	fmt.Printf("Listening on http://localhost%s\n", p)
	log.Fatal(http.ListenAndServe(p, Router))
}

// ServeSSL will start a HTTPS server on port given
// @param port - An int value in range 1 - 65535
// @param certFile - The path to your cert file
// @param keyFile - The path to your key file
func ServeSSL(port int, certFile string, keyFile string) {
	p := ":" + strconv.Itoa(port)
	if port > 65535 || port < 1 {
		fmt.Println(errors.New("port out of range must be in range 1 - 65535"))
	}
	Router.Use(defaultHeaderHandler)
	fmt.Printf("Listening on https://localhost%s\n", p)
	log.Fatal(http.ListenAndServeTLS(p, certFile, keyFile, Router))
}
