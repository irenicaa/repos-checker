package github

import (
	"github.com/irenicaa/repos-checker/loader"
	sourceutils "github.com/irenicaa/repos-checker/loader/sources/source-utils"
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
	repos, err := sourceutils.GetAllPages(func(page int) ([]string, error) {
		return GetReposPage(source.Owner, maxPageSize, page)
	})
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
