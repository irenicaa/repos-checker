package filesystem

import (
	"github.com/irenicaa/repos-checker/loader"
	sourceutils "github.com/irenicaa/repos-checker/loader/sources/source-utils"
	"github.com/irenicaa/repos-checker/models"
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
			if page > 1 {
				return nil, nil
			}

			repos, err := GetRepos(source.BasePath)
			if err == ErrItIsRepo {
				return []string{source.BasePath}, nil
			}

			return repos, err
		},
		GetLastCommit,
		source.Logger,
	)
}
