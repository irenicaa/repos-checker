package gitlab

import (
	"github.com/irenicaa/repos-checker/loader"
	sourceutils "github.com/irenicaa/repos-checker/loader/sources/source-utils"
	"github.com/irenicaa/repos-checker/models"
)

// Source ...
type Source struct {
	Owner    string
	IsGroup  bool
	PageSize int
	Logger   loader.Logger `json:"-"`
}

// Name ...
func (source Source) Name() string {
	return "gitlab:" + source.Owner
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	return sourceutils.LoadRepos(
		source.Name(),
		func(page int) ([]string, error) {
			return GetReposPage(source.Owner, source.IsGroup, source.PageSize, page)
		},
		GetLastCommit,
		source.Logger,
	)
}
