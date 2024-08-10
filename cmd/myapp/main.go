package main

import (
	"database/sql"
	"fmt"
	"github.com/godovasik/amongus/pkg/model"
	"log"

	_ "github.com/lib/pq"
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

func main() {
	users := []model.User{}
	users = append(users, model.User{"1", "qwerty", "asdf", 228, 123})
	fmt.Println(users)

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
	//fmt.Println("Successfully created table")

}
