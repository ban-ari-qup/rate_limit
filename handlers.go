package main

import (
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	apiKey := registerUser(name)
	w.Write([]byte(apiKey))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if CheckIfExists(apiKey) {
		err := Users[apiKey].addRequest()
		if err == nil {
			w.Write([]byte("Good one"))
		} else {
			http.Error(w, "limit", http.StatusTooManyRequests)
		}
	} else {
		http.Error(w, "There is no such User", http.StatusUnauthorized)
	}

}
