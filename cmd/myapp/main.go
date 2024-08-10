package main

import (
	"fmt"
	"github.com/godovasik/amongus/pkg/model"
)

func main() {
	users := []model.User{}
	users = append(users, model.User{"1", "qwerty", "asdf", 228, 123})
	fmt.Println(users)

	connstr := "host=192.168.31.32 port=5432 user=amongus password"
}
