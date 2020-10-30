package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/irenicaa/repos-checker/loader/sources/bitbucket"
	"github.com/irenicaa/repos-checker/loader/sources/github"
)

func main() {
	source := flag.String("source", "bitbucket", "")
	pageSize := flag.Int("pageSize", 100, "")
	page := flag.Int("page", 1, "")
	flag.Parse()

	switch *source {
	case "github":
		source := github.Source{Owner: "irenicaa"}
		reposStates, err := source.LoadRepos()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %+v\n", source.Name(), reposStates)
	case "bitbucket":
		repoState, err := bitbucket.GetReposPage("MartinFelis", *pageSize, *page)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", repoState)
	default:
		log.Fatal("unknown source")
	}
}
