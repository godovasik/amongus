package main

import (
	"database/sql"
	"fmt"
	"github.com/godovasik/amongus/pkg/model"
	"github.com/google/uuid"
	_ "github.com/json-iterator/go"
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
func createUser(db *sql.DB, firstName, lastName string, age int) (sql.Result, error) {
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

func getRows(db *sql.DB, minAge, maxAge int, start, end int64) (*sql.Rows, error) {
	var (
		com  string
		rows *sql.Rows
		err  error
	)

	if start == -1 && minAge == -1 {
		com = `select id, first_name, last_name, age, recording_date from Users`
		rows, err = db.Query(com)
	} else if start == -1 {
		com = `select id, first_name, last_name, age, recording_date from Users
				where age between $1 and $2`
		rows, err = db.Query(com, minAge, maxAge)
	} else if minAge == -1 {
		com = `select id, first_name, last_name, age, recording_date from Users
				where recording_date between $1 and $2`
		rows, err = db.Query(com, start, end)

	} else {
		com = `select id, first_name, last_name, age, recording_date from Users
    			where age between $1 and $2 and recording_date between $3 and $4`
		rows, err = db.Query(com, minAge, maxAge, start, end)
	}
	return rows, err
}

func getUsersFromRange(db *sql.DB, minAge, maxAge int, start, end int64) ([]model.User, error) {
	rows, err := getRows(db, minAge, maxAge, start, end)
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

func initDB() (*sql.DB, error) {
	connStr := "host=185.221.162.204 port=5432 user=lesha password=amongus dbname=test sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	//fmt.Println("kek")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	err = createTable(db)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return nil, err
	}
	return db, nil
}

func printUsers(users []model.User) {
	for _, user := range users {
		fmt.Println(user.FirstName, user.LastName, user.Age, user.RecordingDate)
	}
}

func main() {
	db, err := initDB()
	if err != nil {
		return
	}
	defer db.Close()

	//users := []model.User{}

	//users = append(users, model.User{uuid.New().String(), "Alex", "Joshson", 5, 100})
	//users = append(users, model.User{uuid.New().String(), "Bob", "Marley", 11, 150})
	//users = append(users, model.User{uuid.New().String(), "C", "B", 26, 200})
	//users = append(users, model.User{uuid.New().String(), "Alex", "Joshson", 45, 300})
	//users = append(users, model.User{uuid.New().String(), "k", "ek", 90, 500})

	//result, err := createUser(db, )
	//if err != nil {
	//	fmt.Println("Error creating user:", err)
	//}

	//fmt.Println(result.RowsAffected())

	//users, _ := getUsers(db)

	users, _ := getUsersFromRange(db, 27, 89, -1, 0)
	printUsers(users)

	//_, err = db.Exec("DROP TABLE users")
}
