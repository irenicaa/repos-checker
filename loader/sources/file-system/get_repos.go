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

	var allRepos []string
	for _, fileInfo := range filesInfos {
		if !fileInfo.IsDir() {
			continue
		}
		if fileInfo.Name() == ".git" {
			return nil, ErrItIsRepo
		}

		directoryPath := filepath.Join(basePath, fileInfo.Name())
		repos, err := GetRepos(directoryPath)
		switch err {
		case nil:
		case ErrItIsRepo:
			allRepos = append(allRepos, directoryPath)
		default:
			return nil, fmt.Errorf("unable to get repos from the subdirectory: %v", err)
		}

		allRepos = append(allRepos, repos...)
	}

	return allRepos, nil
}
