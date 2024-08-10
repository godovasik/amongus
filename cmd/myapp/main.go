package main

import (
	"database/sql"
	"fmt"
	"github.com/godovasik/amongus/pkg/model"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	users := []model.User{}
	users = append(users, model.User{"1", "qwerty", "asdf", 228, 123})
	fmt.Println(users)

	connStr := "host=185.221.162.204 port=5432 user=lesha password=amongus dbname=test sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	fmt.Println("kek")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
