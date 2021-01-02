package main

import (
	"encoding/json"
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
	"github.com/irenicaa/repos-checker/models"
)

func main() {
	source := flag.String("source", "", "")
	owner := flag.String("owner", "", "")
	isGroup := flag.Bool("group", false, "")
	pageSize := flag.Int("pageSize", 100, "")
	basePath := flag.String("path", "..", "")
	command := flag.String("command", "./tools/test_tool.bash ..", "")
	flag.Parse()
	if *source == "" {
		log.Fatal("source is unspecified")
	}

	var sourceInstance loader.Source
	logger := log.New(os.Stderr, "", log.LstdFlags)
	switch *source {
	case "github":
		if *owner == "" {
			*owner = "irenicaa"
		}

		sourceInstance = github.Source{
			Owner:    *owner,
			PageSize: *pageSize,
			Logger:   logger,
		}
	case "bitbucket":
		if *owner == "" {
			*owner = "MartinFelis"
		}

		sourceInstance = bitbucket.Source{
			Workspace: *owner,
			PageSize:  *pageSize,
			Logger:    logger,
		}
	case "gitlab":
		if *owner == "" {
			*owner = "dzaporozhets"
		}

		sourceInstance = gitlab.Source{
			Owner:    *owner,
			IsGroup:  *isGroup,
			PageSize: *pageSize,
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
			Command:   commandParts[0],
			Arguments: commandParts[1:],
		}
	default:
		log.Fatal("unknown source")
	}

	sourceState, err := loader.LoadSource(sourceInstance)
	if err != nil {
		log.Fatalf("unable to load repos: %s", err)
	}

	duplicates := models.FindRepoStateDuplicates(sourceState.Repos)
	if len(duplicates) != 0 {
		log.Printf("repos has duplicates: %v", duplicates)
	}

	sourceStateBytes, err := json.Marshal(sourceState)
	if err != nil {
		log.Fatalf("unable to marshal the source state: %s", err)
	}

	fmt.Println(string(sourceStateBytes))
}
