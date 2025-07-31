package manage

import (
	"errors"
	"mycity/manage/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID   string
	Name string
}

func GetUser(token string) (User, error) {
	// Placeholder for user retrieval logic
	// In a real implementation, this would query the database
	rows, err := database.GetUser(token)
	if err != nil {
		return User{}, err
	}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err == nil {
			return user, nil
		}
	}
	return User{}, errors.New("length is 0")
}

func User_RegisterRoutes(r chi.Router) {
	r.Get("/preference", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
}
