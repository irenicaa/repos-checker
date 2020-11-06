package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/irenicaa/repos-checker/loader/sources/bitbucket"
	filesystem "github.com/irenicaa/repos-checker/loader/sources/file-system"
	"github.com/irenicaa/repos-checker/loader/sources/github"
	"github.com/irenicaa/repos-checker/loader/sources/gitlab"
)

func main() {
	source := flag.String("source", "file-system", "")
	flag.Parse()

	logger := log.New(os.Stderr, "", log.LstdFlags)
	switch *source {
	case "github":
		source := github.Source{Owner: "irenicaa", Logger: logger}
		reposStates, err := source.LoadRepos()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %+v\n", source.Name(), reposStates)
	case "bitbucket":
		source := bitbucket.Source{Workspace: "MartinFelis", Logger: logger}
		reposStates, err := source.LoadRepos()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %+v\n", source.Name(), reposStates)
	case "gitlab":
		source := gitlab.Source{Owner: "dzaporozhets", Logger: logger}
		reposStates, err := source.LoadRepos()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %+v\n", source.Name(), reposStates)
	case "file-system":
		source := filesystem.Source{BasePath: ".."}
		reposStates, err := source.LoadRepos()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %+v\n", source.Name(), reposStates)
	default:
		log.Fatal("unknown source")
	}
}
