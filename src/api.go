package main

import (
	"net/http"
)

// HandleStats ...
func HandleStats(w http.ResponseWriter, r *http.Request) {

}

// HandleHealth ...
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
