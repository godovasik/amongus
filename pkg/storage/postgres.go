package storage

import (
	"database/sql"
	"fmt"
	"github.com/godovasik/amongus/pkg/model"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

func CreateTable(db *sql.DB) error {
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
func CreateUser(db *sql.DB, firstName, lastName string, age int) (sql.Result, error) {
	userID := uuid.New().String()
	timeStamp := time.Now().UnixMilli()
	user := model.User{userID, firstName, lastName, age, timeStamp}

	com := `
        insert into Users (ID, first_name, last_name, age, recording_date)
        values ($1, $2, $3, $4, $5)
    `
	result, err := db.Exec(com, user.ID, user.FirstName, user.LastName, user.Age, user.RecordingDate)
	return result, err
}

func GetUsers(db *sql.DB) ([]model.User, error) {
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

func GetQuery(minAge, maxAge *int, start, end *int64) (string, []any) {

	query := "select * from users where 1 = 1"
	var params []any
	paramCount := 1

	if minAge != nil {
		query += fmt.Sprintf(" and age >= $%d", paramCount)
		params = append(params, *minAge)
		paramCount++
	}

	if maxAge != nil {
		query += fmt.Sprintf(" and age <= $%d", paramCount)
		params = append(params, *maxAge)
		paramCount++
	}

	if start != nil {
		query += fmt.Sprintf(" and recording_date >= $%d", paramCount)
		params = append(params, *start)
		paramCount++
	}

	if end != nil {
		query += fmt.Sprintf(" and recording_date <= $%d", paramCount)
		params = append(params, *end)
	}

	return query, params
}

func GetUsersFromRange(db *sql.DB, minAgeP, maxAgeP *int, startP, endP *int64) ([]model.User, error) {
	query, params := GetQuery(minAgeP, maxAgeP, startP, endP)

	//fmt.Println(query)

	rows, err := db.Query(query, params...)
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

func InitDB() (*sql.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	//fmt.Println("kek")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	err = CreateTable(db)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return nil, err
	}
	return db, nil
}

func PrintUsers(users []model.User) {
	for _, user := range users {
		fmt.Println(user.FirstName, user.LastName, user.Age, user.RecordingDate)
	}
}
