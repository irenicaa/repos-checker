package bitbucket

import (
	"github.com/irenicaa/repos-checker/v2/loader"
	sourceutils "github.com/irenicaa/repos-checker/v2/loader/sources/source-utils"
	"github.com/irenicaa/repos-checker/v2/models"
)

// Source ...
type Source struct {
	Workspace string
	PageSize  int
	Logger    loader.Logger `json:"-"`
}

// Name ...
func (source Source) Name() string {
	return "bitbucket:" + source.Workspace
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	return sourceutils.LoadRepos(
		source.Name(),
		func(page int) ([]string, error) {
			return GetReposPage(source.Workspace, source.PageSize, page)
		},
		func(repo string) (models.RepoState, error) {
			return GetLastCommit(source.Workspace, repo)
		},
		source.Logger,
	)
}
