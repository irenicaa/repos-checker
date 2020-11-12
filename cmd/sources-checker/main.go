package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/irenicaa/repos-checker/loader/sources/bitbucket"
	"github.com/irenicaa/repos-checker/loader/sources/external"
	filesystem "github.com/irenicaa/repos-checker/loader/sources/file-system"
	"github.com/irenicaa/repos-checker/loader/sources/github"
	"github.com/irenicaa/repos-checker/loader/sources/gitlab"
	"github.com/irenicaa/repos-checker/models"
)

func main() {
	source := flag.String("source", "external", "")
	flag.Parse()

	const maxPageSize = 100
	logger := log.New(os.Stderr, "", log.LstdFlags)
	var sourceName string
	var reposStates []models.RepoState
	var err error
	switch *source {
	case "github":
		source := github.Source{
			Owner:       "irenicaa",
			MaxPageSize: maxPageSize,
			Logger:      logger,
		}
		sourceName = source.Name()
		reposStates, err = source.LoadRepos()
	case "bitbucket":
		source := bitbucket.Source{
			Workspace:   "MartinFelis",
			MaxPageSize: maxPageSize,
			Logger:      logger,
		}
		sourceName = source.Name()
		reposStates, err = source.LoadRepos()
	case "gitlab":
		source := gitlab.Source{
			Owner:       "dzaporozhets",
			MaxPageSize: maxPageSize,
			Logger:      logger,
		}
		sourceName = source.Name()
		reposStates, err = source.LoadRepos()
	case "file-system":
		source := filesystem.Source{BasePath: "..", Logger: logger}
		sourceName = source.Name()
		reposStates, err = source.LoadRepos()
	case "external":
		source := external.Source{
			AdditionalName: "test",
			Command:        "./tools/test_tool.bash",
			Arguments:      []string{".."},
		}
		sourceName = source.Name()
		reposStates, err = source.LoadRepos()
	default:
		log.Fatal("unknown source")
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %+v\n", sourceName, reposStates)
}
