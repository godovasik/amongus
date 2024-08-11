package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/godovasik/amongus/pkg/storage"
	"net/http"
	"strconv"
	"time"
)

func NewUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//if r.Method != "POST" {
		//	http.Error(w, "Post Only", http.StatusMethodNotAllowed)
		//	return
		//}

		if r.URL.Path != "/createUser" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		age := r.URL.Query().Get("age")
		if len(age) == 0 {
			http.Error(w, "age is required", http.StatusBadRequest)
			return
		}

		ageInt, err := strconv.Atoi(age)
		if err != nil {
			http.Error(w, "invalid age", http.StatusBadRequest)
			return
		}

		firstname := r.URL.Query().Get("firstname")
		if len(firstname) == 0 {
			http.Error(w, "firstname required", http.StatusBadRequest)
			return
		}

		lastname := r.URL.Query().Get("lastname")
		if len(lastname) == 0 {
			http.Error(w, "lastname required", http.StatusBadRequest)
			return
		}

		_, err = storage.CreateUser(db, firstname, lastname, ageInt)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = w.Write([]byte("GOOOOL"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}

func ListUsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Get Only", http.StatusMethodNotAllowed)
		}

		users, err := storage.GetUsers(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		usersJson, _ := json.Marshal(users)

		http.ServeContent(w, r, "users.json", time.Now(), bytes.NewReader(usersJson))
		//w.Write(usersJson)
		return
	}
}

func getParamsForRange(r *http.Request) (minAge, maxAge *int, start, end *int64, err error) {
	minAgeStr := r.URL.Query().Get("minAge")
	minAge = new(int)
	if minAgeStr == "" {
		minAge = nil
	} else {
		*minAge, err = strconv.Atoi(minAgeStr)
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf("invalid minAge: %s", minAgeStr)
		}
	}

	maxAgeStr := r.URL.Query().Get("maxAge")
	maxAge = new(int)
	if maxAgeStr == "" {
		maxAge = nil
	} else {
		*maxAge, err = strconv.Atoi(maxAgeStr)
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf("invalid maxAge: %s", maxAgeStr)
		}
	}

	startStr := r.URL.Query().Get("start")
	start = new(int64)
	if startStr == "" {
		start = nil
	} else {
		*start, err = strconv.ParseInt(startStr, 10, 64)
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf("invalid start: %s", startStr)
		}
	}

	endStr := r.URL.Query().Get("end")
	end = new(int64)
	if endStr == "" {
		end = nil
	} else {
		*end, err = strconv.ParseInt(endStr, 10, 64)
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf("invalid end: %s", endStr)
		}
	}

	return minAge, maxAge, start, end, nil
}

func ListUsersFromRangeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Get Only", http.StatusMethodNotAllowed)
			return
		}

		minAgeP, maxAgeP, startP, endP, err := getParamsForRange(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		users, err := storage.GetUsersFromRange(db, minAgeP, maxAgeP, startP, endP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		usersJson, _ := json.Marshal(users)
		http.ServeContent(w, r, "users.json", time.Now(), bytes.NewReader(usersJson))
	}
}
