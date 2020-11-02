package bitbucket

import (
	"github.com/irenicaa/repos-checker/loader"
	"github.com/irenicaa/repos-checker/models"
)

// Source ...
type Source struct {
	Workspace string
	Logger    loader.Logger
}

// Name ...
func (source Source) Name() string {
	return "bitbucket"
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	const maxPageSize = 100
	repos, err := GetRepos(source.Workspace, maxPageSize)
	if err != nil {
		return nil, err
	}

	var reposStates []models.RepoState
	for _, repo := range repos {
		repoState, err := GetLastCommit(source.Workspace, repo)
		switch err {
		case nil:
		case errNoCommits:
			source.Logger.Printf(
				"%s repo that belongs to %s has no commits",
				repo,
				source.Workspace,
			)
		default:
			return nil, err
		}

		reposStates = append(reposStates, repoState)
	}

	return reposStates, nil
}
