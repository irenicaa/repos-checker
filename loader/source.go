package loader

import "github.com/irenicaa/repos-checker/models"

// Source ...
type Source interface {
	Name() string
	LoadRepos() ([]models.RepoState, error)
}
