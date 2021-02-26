package grapi

import "net/http"

// ResponseStruct is a struct for to help prevent code duplication
type ResponseStruct struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// Header header struct for adding headers
type Header struct {
	Key string
	Value string
}

var defaultHeaders []Header = []Header{};

// Get will setup a GET handle function
// @param path - The path the server will check on request
// @param f - A function to tell it what to do on request
func Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	Router.HandleFunc(path, f).Methods("GET")
}

// Post will setup a POST handle function
// @param path - The path the server will check on request
// @param f - A function to tell it what to do on request
func Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	Router.HandleFunc(path, f).Methods("POST")
}

// Delete will setup a DELETE handle function
// @param path - The path the server will check on request
// @param f - A function to tell it what to do on request
func Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	Router.HandleFunc(path, f).Methods("DELETE")
}

// Patch will setup a PATCH handle function
// @param path - The path the server will check on request
// @param f - A function to tell it what to do on request
func Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
	Router.HandleFunc(path, f).Methods("PATCH")
}

// AddDefaultHeader will add the header given in every request
// @param key - The key value of header
// @param value - The value of the key
func AddDefaultHeader(key string, value string) {
	defaultHeaders = append(defaultHeaders, Header{
		key,
		value,
	})
}

// AddDefaultHeaders will add headers with given list
// @param headers - A list of type Header
func AddDefaultHeaders(headers []Header) {
	for _, h := range headers {
		defaultHeaders = append(defaultHeaders, Header{
			h.Key,
			h.Value,
		})
	}
}

// defaultHeaderHandler will add the headers given by AddDefaultHeaders/AddDefaultHeader
func defaultHeaderHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range defaultHeaders {
			w.Header().Add(h.Key, h.Value)
		} 
		next.ServeHTTP(w, r)
	})
}
