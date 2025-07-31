package main

import (
	"fmt"
	"log"
	"mycity/manage"
	"mycity/manage/database"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func TokenSetter(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:  "token",
		Value: token,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
	database.CreateUser(token, "Default User")
	fmt.Println("New user detected, created with token:", token)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil || token.Value == "" {
			TokenSetter(w, uuid.New().String())
		} else {
			_, err := manage.GetUser(token.Value) // Ensure user exists
			if err != nil {
				TokenSetter(w, token.Value)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not set
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(AuthMiddleware)

	fileServer := http.FileServer(http.Dir("./static/"))
	r.Route("/api", func(r2 chi.Router) {
		manage.Map_RegisterRoutes(r2)
		manage.User_RegisterRoutes(r2)
	})
	database.InitDatabase()
	r.Handle("/*", http.StripPrefix("/", fileServer))

	fmt.Println("Starting server on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
