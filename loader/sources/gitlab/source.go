package gitlab

import (
	"github.com/irenicaa/repos-checker/loader"
	sourceutils "github.com/irenicaa/repos-checker/loader/sources/source-utils"
	"github.com/irenicaa/repos-checker/models"
)

// Source ...
type Source struct {
	Owner    string
	PageSize int
	Logger   loader.Logger
}

// Name ...
func (source Source) Name() string {
	return "gitlab"
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	return sourceutils.LoadRepos(
		func(page int) ([]string, error) {
			return GetReposPage(source.Owner, source.PageSize, page)
		},
		GetLastCommit,
		source.Logger,
	)
}
