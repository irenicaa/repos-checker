package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	reposchecker "github.com/irenicaa/repos-checker"
	"github.com/irenicaa/repos-checker/config"
)

func main() {
	configPath := flag.String("config", "config.json", "")
	flag.Parse()

	configFile, err := os.OpenFile(*configPath, os.O_RDONLY, 0)
	if err != nil {
		log.Fatalf("unable to open the config file: %v", err)
	}
	defer configFile.Close()

	logger := log.New(os.Stderr, "", log.LstdFlags)
	sources, referenceName, err := config.LoadConfig(configFile, logger)
	if err != nil {
		log.Fatalf("unable to load the config: %v", err)
	}
	if referenceName == "" {
		log.Fatalf("reference isn't specified")
	}

	sourceDiffs := reposchecker.CheckSources(sources, referenceName, logger)
	sourceDiffsBytes, err := json.Marshal(sourceDiffs)
	if err != nil {
		log.Fatalf("unable to marshal source diffs: %s", err)
	}

	fmt.Println(string(sourceDiffsBytes))
}
