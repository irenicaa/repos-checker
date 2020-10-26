package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/irenicaa/repos-checker/loader/sources/github"
)

func main() {
	pageSize := flag.Int("pageSize", 100, "")
	page := flag.Int("page", 1, "")
	flag.Parse()

	repos, err := github.GetReposPage("irenicaa", *pageSize, *page)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", repos)
}
