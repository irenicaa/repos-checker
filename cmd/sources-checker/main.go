package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/irenicaa/repos-checker/loader/sources/github"
)

func main() {
	pageSize := flag.Int("pageSize", 100, "")
	flag.Parse()

	repos, err := github.GetRepos("irenicaa", *pageSize)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", repos)
}
