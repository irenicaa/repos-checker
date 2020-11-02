package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/irenicaa/repos-checker/loader/sources/bitbucket"
	"github.com/irenicaa/repos-checker/loader/sources/github"
	"github.com/irenicaa/repos-checker/loader/sources/gitlab"
)

func main() {
	source := flag.String("source", "gitlab", "")
	pageSize := flag.Int("pageSize", 100, "")
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
		source := bitbucket.Source{Workspace: "MartinFelis"}
		reposStates, err := source.LoadRepos()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %+v\n", source.Name(), reposStates)
	case "gitlab":
		reposPaths, err := gitlab.GetRepos("dzaporozhets", *pageSize)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", reposPaths)
	default:
		log.Fatal("unknown source")
	}
}
