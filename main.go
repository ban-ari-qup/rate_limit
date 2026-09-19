package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/register", registerHandler)
	r.Get("/hello", helloHandler)

	http.ListenAndServe(":3000", r)
}
