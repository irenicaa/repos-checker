package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/irenicaa/repos-checker/loader/sources/bitbucket"
	"github.com/irenicaa/repos-checker/loader/sources/github"
)

type workspace struct {
	Name string
	Slug string
	Type string
}

func main() {
	source := flag.String("source", "bitbucket", "")
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
		var workspace workspace // nolint: vetshadow
		if err := bitbucket.SendRequest("/workspaces/MartinFelis", nil, &workspace); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", workspace)
	default:
		log.Fatal("unknown source")
	}
}
