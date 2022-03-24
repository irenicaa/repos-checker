package loader

import "github.com/irenicaa/repos-checker/v2/models"

// Source ...
type Source interface {
	Name() string
	LoadRepos() ([]models.RepoState, error)
}
