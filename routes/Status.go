package routes

import (
	"net/http"
	"encoding/json"
) 

//StatusRoute ..
type StatusRoute struct {
	StatusCode int `json:"statusCode"`
}

//Status ...
func Status(w http.ResponseWriter, r *http.Request) {
	d := StatusRoute{
		http.StatusOK,
	}
	json.NewEncoder(w).Encode(d)
}