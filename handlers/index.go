package handlers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json; charset=UTF=8")
	w.WriteHeader(http.StatusOK)
}

