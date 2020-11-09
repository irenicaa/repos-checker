package bitbucket

import (
	"github.com/irenicaa/repos-checker/loader"
	sourceutils "github.com/irenicaa/repos-checker/loader/sources/source-utils"
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
	repos, err := sourceutils.GetAllPages(func(page int) ([]string, error) {
		return GetReposPage(source.Workspace, maxPageSize, page)
	})
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
