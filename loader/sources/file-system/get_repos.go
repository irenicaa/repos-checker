package filesystem

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ErrItIsRepo ...
var ErrItIsRepo = errors.New("it is repo")

// GetRepos ...
func GetRepos(basePath string) ([]string, error) {
	file, err := os.OpenFile(basePath, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("unable to open the base path: %v", err)
	}
	defer file.Close()

	filesInfos, err := file.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("unable to read a subdirectory list: %v", err)
	}

	var allReposPaths []string
	for _, fileInfo := range filesInfos {
		if !fileInfo.IsDir() {
			continue
		}
		if fileInfo.Name() == ".git" {
			return nil, ErrItIsRepo
		}

		directoryPath := filepath.Join(basePath, fileInfo.Name())
		reposPaths, err := GetRepos(directoryPath)
		switch err {
		case nil:
		case ErrItIsRepo:
			allReposPaths = append(allReposPaths, directoryPath)
		default:
			return nil, fmt.Errorf(
				"unable to get repos paths from the subdirectory %s: %v",
				directoryPath,
				err,
			)
		}

		allReposPaths = append(allReposPaths, reposPaths...)
	}

	return allReposPaths, nil
}
