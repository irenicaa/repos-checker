package gitlab

import (
	"github.com/irenicaa/repos-checker/loader"
	"github.com/irenicaa/repos-checker/models"
)

// Source ...
type Source struct {
	Owner  string
	Logger loader.Logger
}

// Name ...
func (source Source) Name() string {
	return "gitlab"
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	const maxPageSize = 100
	reposPaths, err := GetRepos(source.Owner, maxPageSize)
	if err != nil {
		return nil, err
	}

	var reposStates []models.RepoState
	for _, repoPath := range reposPaths {
		repoState, err := GetLastCommit(repoPath)
		switch err {
		case nil:
		case errNoCommits:
			source.Logger.Printf("%s repo has no commits", repoPath)
		default:
			return nil, err
		}

		reposStates = append(reposStates, repoState)
	}

	return reposStates, nil
}
