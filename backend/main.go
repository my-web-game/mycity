package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/*", http.StripPrefix("/", fileServer))

	http.ListenAndServe(":3000", r)
}
