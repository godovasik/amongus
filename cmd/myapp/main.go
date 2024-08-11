package main

import (
	_ "github.com/godovasik/amongus/pkg/api"
	"github.com/godovasik/amongus/pkg/storage"
	_ "github.com/json-iterator/go"
	_ "github.com/lib/pq"
)

func main() {

	db, err := storage.InitDB()
	if err != nil {
		return
	}
	defer db.Close()

	//http.HandleFunc("/hello", api.NewUserHandler(db))
	//
	//fmt.Println("Pognali naxou")
	//
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}

	users, _ := storage.GetUsers(db)
	storage.PrintUsers(users)

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

	//users, _ := getUsersFromRange(db, 27, 89, -1, 0)
	//printUsers(users)

	//_, err = db.Exec("DROP TABLE users")
}
