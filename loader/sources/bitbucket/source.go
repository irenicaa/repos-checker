package bitbucket

import (
	"github.com/irenicaa/repos-checker/loader"
	sourceutils "github.com/irenicaa/repos-checker/loader/sources/source-utils"
	"github.com/irenicaa/repos-checker/models"
)

// Source ...
type Source struct {
	Workspace string
	PageSize  int
	Logger    loader.Logger
}

// Name ...
func (source Source) Name() string {
	return "bitbucket"
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	return sourceutils.LoadRepos(
		func(page int) ([]string, error) {
			return GetReposPage(source.Workspace, source.PageSize, page)
		},
		func(repo string) (models.RepoState, error) {
			return GetLastCommit(source.Workspace, repo)
		},
		source.Logger,
	)
}
