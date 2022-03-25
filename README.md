# Repos Checker

[![GoDoc](https://godoc.org/github.com/irenicaa/repos-checker/v2?status.svg)](https://godoc.org/github.com/irenicaa/repos-checker/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/irenicaa/repos-checker/v2)](https://goreportcard.com/report/github.com/irenicaa/repos-checker/v2)
[![Build Status](https://app.travis-ci.com/irenicaa/repos-checker.svg?branch=master)](https://app.travis-ci.com/irenicaa/repos-checker)
[![codecov](https://codecov.io/gh/irenicaa/repos-checker/branch/master/graph/badge.svg)](https://codecov.io/gh/irenicaa/repos-checker)

The library and the utilities for checking that repo mirrors are up to date. Checking is done by comparing the latest commits in each mirror of the repo.

Cloud services (GitHub, GitLab, and Bitbucket), local copies, and any external program that returns states of repos are supported as sources of mirrors. Also, any such sources can be combined into one with merging of the lists of the repos they contain.

## Installation

```
$ go get github.com/irenicaa/repos-checker/...
```

## Tools

- [repos-checker](cmd/repos-checker) &mdash; utility for checking that repo mirrors are up to date
- [sources-checker](cmd/sources-checker) &mdash; utility for getting a list of the latest commits of repos from a specific source

## Docs

- [config.schema.json](docs/config.schema.json) &mdash; JSON Schema for config of sources
- [config.example.json](docs/config.example.json) &mdash; example for config of sources

## License

The MIT License (MIT)

Copyright &copy; 2021 irenica
