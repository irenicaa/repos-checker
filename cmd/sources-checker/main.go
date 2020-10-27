package main

import (
	"fmt"
	"log"

	"github.com/irenicaa/repos-checker/loader/sources/github"
)

func main() {
	source := github.Source{Owner: "irenicaa"}
	reposStates, err := source.LoadRepos()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %+v\n", source.Name(), reposStates)
}
