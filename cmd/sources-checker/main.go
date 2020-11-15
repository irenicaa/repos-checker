package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/irenicaa/repos-checker/loader"
	"github.com/irenicaa/repos-checker/loader/sources/bitbucket"
	"github.com/irenicaa/repos-checker/loader/sources/external"
	filesystem "github.com/irenicaa/repos-checker/loader/sources/file-system"
	"github.com/irenicaa/repos-checker/loader/sources/github"
	"github.com/irenicaa/repos-checker/loader/sources/gitlab"
)

func main() {
	source := flag.String("source", "", "")
	owner := flag.String("owner", "", "")
	basePath := flag.String("path", "..", "")
	command := flag.String("command", "./tools/test_tool.bash ..", "")
	flag.Parse()
	if *source == "" {
		log.Fatal("source is unspecified")
	}

	const maxPageSize = 100
	logger := log.New(os.Stderr, "", log.LstdFlags)
	var sourceInstance loader.Source
	switch *source {
	case "github":
		if *owner == "" {
			*owner = "irenicaa"
		}

		sourceInstance = github.Source{
			Owner:    *owner,
			PageSize: maxPageSize,
			Logger:   logger,
		}
	case "bitbucket":
		if *owner == "" {
			*owner = "MartinFelis"
		}

		sourceInstance = bitbucket.Source{
			Workspace: *owner,
			PageSize:  maxPageSize,
			Logger:    logger,
		}
	case "gitlab":
		if *owner == "" {
			*owner = "dzaporozhets"
		}

		sourceInstance = gitlab.Source{
			Owner:    *owner,
			PageSize: maxPageSize,
			Logger:   logger,
		}
	case "file-system":
		sourceInstance = filesystem.Source{BasePath: *basePath, Logger: logger}
	case "external":
		commandParts := strings.Fields(*command)
		if len(commandParts) == 0 {
			log.Fatal("command shouldn't be empty")
		}

		sourceInstance = external.Source{
			AdditionalName: "test",
			Command:        commandParts[0],
			Arguments:      commandParts[1:],
		}
	default:
		log.Fatal("unknown source")
	}

	reposStates, err := sourceInstance.LoadRepos()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %+v\n", sourceInstance.Name(), reposStates)
}
