package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/irenicaa/repos-checker/loader/sources/github"
)

func main() {
	source := flag.String("source", "github", "")
	flag.Parse()

	switch *source {
	case "github":
		source := github.Source{Owner: "irenicaa"}
		reposStates, err := source.LoadRepos()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %+v\n", source.Name(), reposStates)
	default:
		log.Fatal("unknown source")
	}
}
