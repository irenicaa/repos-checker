package main

import (
	"fmt"
	"log"

	"github.com/irenicaa/repos-checker/loader/sources/github"
)

// User ...
type User struct {
	Name      string
	Followers int
}

func main() {
	var user User
	if err := github.SendRequest("/users/irenicaa", nil, &user); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v	\n", user)
}
