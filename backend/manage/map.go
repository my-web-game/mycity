package manage

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type MapHeader struct {
	ID          string
	Name        string
	Description string
	Thumbnail   string
	CreatedAt   time.Time
	UupdatedAt  time.Time
}

func GetMapCount(token string) (int, error) {
	return 0, nil
}

func GetMap(index int) (MapHeader, error) {
	return MapHeader{}, errors.New("not implemented")
}

func Map_RegisterRoutes(r chi.Router) {
	r.Get("/maps", func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		count, err := GetMapCount(token.Value)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(strconv.Itoa(count)))
	})
}
