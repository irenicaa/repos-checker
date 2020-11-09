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
	return sourceutils.LoadRepos(
		func(page int) ([]string, error) {
			return GetReposPage(source.Owner, maxPageSize, page)
		},
		func(repo string) (models.RepoState, error) {
			return GetLastCommit(source.Owner, repo)
		},
		source.Logger,
	)
}
