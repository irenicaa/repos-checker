package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/irenicaa/repos-checker/loader/sources/bitbucket"
	"github.com/irenicaa/repos-checker/loader/sources/github"
	"github.com/irenicaa/repos-checker/loader/sources/gitlab"
)

type user struct {
	Name  string
	State string
}

func main() {
	source := flag.String("source", "gitlab", "")
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
		parameters := url.Values{}
		parameters.Add("username", "dzaporozhets")

		var users []user
		err := gitlab.SendRequest("/users", parameters, &users)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", users)
	default:
		log.Fatal("unknown source")
	}
}
