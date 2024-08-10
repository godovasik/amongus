package main

import (
	"database/sql"
	"fmt"
	"github.com/godovasik/amongus/pkg/model"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func createTable(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
   id VARCHAR(255),
   first_name VARCHAR(255),
   last_name VARCHAR(255),
   age INTEGER,
   recording_date BIGINT
  );
    `)
	return err
}
func createUser(db *sql.DB, user model.User) (sql.Result, error) {

	com := `
        insert into Users (ID, first_name, last_name, age, recording_date)
        values ($1, $2, $3, $4, $5)
    `
	result, err := db.Exec(com, user.ID, user.FirstName, user.LastName, user.Age, user.RecordingDate)
	return result, err
}

func getUsers(db *sql.DB) ([]model.User, error) {
	com := `
       select id, first_name, last_name, age, recording_date from Users
   `
	rows, err := db.Query(com)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.RecordingDate)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, err
}

func getCom(start int64, minAge int) string {
	var com string
	if start == -1 && minAge == -1 {
		com = `select id, first_name, last_name, age, recording_date from Users`
	} else if start == -1 {
		com = `select id, first_name, last_name, age, recording_date from Users
				where age between $1 and $2`
	} else if minAge == -1 {
		com = `select id, first_name, last_name, age, recording_date from Users
				where recording_date between $3 and $4`
	} else {
		com = `select id, first_name, last_name, age, recording_date from Users
    			where age between $1 and $2 and recording_date between $3 and $4`
	}
	return com
}

func getUsersFromRange(db *sql.DB, start, end int64, minAge, maxAge int) ([]model.User, error) {
	com := getCom(start, minAge)
	rows, err := db.Query(com, minAge, maxAge, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.RecordingDate)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, err
}

func main() {
	userID := uuid.New().String()
	timeStamp := time.Now().UnixMilli()
	user := model.User{userID, "qwerty", "asdf", 228, timeStamp}
	fmt.Println(user)

	connStr := "host=185.221.162.204 port=5432 user=lesha password=amongus dbname=test sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	//fmt.Println("kek")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	err = createTable(db)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	result, err := createUser(db, user)
	if err != nil {
		fmt.Println("Error creating user:", err)
	}

	fmt.Println(result.RowsAffected())

	fmt.Println(getUsers(db))
	_, err = db.Exec("DROP TABLE my_table")
}
