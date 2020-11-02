package github

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
	return "github"
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	const maxPageSize = 100
	repos, err := GetRepos(source.Owner, maxPageSize)
	if err != nil {
		return nil, err
	}

	var reposStates []models.RepoState
	for _, repo := range repos {
		repoState, err := GetLastCommit(source.Owner, repo)
		switch err {
		case nil:
		case errNoCommits:
			source.Logger.Printf(
				"%s repo that belongs to %s has no commits",
				repo,
				source.Owner,
			)
		default:
			return nil, err
		}

		reposStates = append(reposStates, repoState)
	}

	return reposStates, nil
}
