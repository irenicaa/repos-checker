package filesystem

import (
	"github.com/irenicaa/repos-checker/models"
)

// Source ...
type Source struct {
	BasePath string
}

// Name ...
func (source Source) Name() string {
	return "file-system"
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	reposPaths, err := GetRepos(source.BasePath)
	if err != nil {
		return nil, err
	}

	var reposStates []models.RepoState
	for _, repoPath := range reposPaths {
		repoState, err := GetLastCommit(repoPath)
		if err != nil {
			return nil, err
		}

		reposStates = append(reposStates, repoState)
	}

	return reposStates, nil
}
