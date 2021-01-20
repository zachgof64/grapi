package routes

import (
	"net/http"
	"encoding/json"
) 


//StatusRoute - Struct for response
type StatusRoute struct {
	StatusCode int `json:"statusCode"`
}

//Status - Sends status to verify api is working
func Status(w http.ResponseWriter, r *http.Request) {
	d := StatusRoute{
		http.StatusOK,
	}
	json.NewEncoder(w).Encode(d)
}