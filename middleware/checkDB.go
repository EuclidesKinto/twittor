package middleware

import (
	"net/http"
	"twittor/db"
	// "github.com/ptilotta/twittor/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexão perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
