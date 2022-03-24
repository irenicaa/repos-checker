package filesystem

import (
	"github.com/irenicaa/repos-checker/v2/loader"
	sourceutils "github.com/irenicaa/repos-checker/v2/loader/sources/source-utils"
	"github.com/irenicaa/repos-checker/v2/models"
)

// Source ...
type Source struct {
	BasePath string
	Logger   loader.Logger `json:"-"`
}

// Name ...
func (source Source) Name() string {
	return "file-system:" + source.BasePath
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	return sourceutils.LoadRepos(
		source.Name(),
		func(page int) ([]string, error) {
			return GetReposPage(source.BasePath, page)
		},
		GetLastCommit,
		source.Logger,
	)
}
