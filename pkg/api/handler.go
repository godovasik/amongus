package api

import (
	"database/sql"
	"github.com/godovasik/amongus/pkg/model"
	"net/http"
)

func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		user.ID = "fuck you"
	}
}
