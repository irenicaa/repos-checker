package main

import (
	"fmt"
	"log"

	"github.com/irenicaa/repos-checker/loader/sources/github"
)

func main() {
	repoState, err := github.GetLastCommit("irenicaa", "repos-checker")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", repoState)
}
