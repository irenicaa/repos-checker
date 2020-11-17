package loader

import (
	"fmt"
	"strings"

	"github.com/irenicaa/repos-checker/models"
)

// MultiSource ...
type MultiSource []Source

// Name ...
func (sources MultiSource) Name() string {
	names := []string{}
	for _, source := range sources {
		name := source.Name()
		names = append(names, name)
	}

	return strings.Join(names, "|")
}

// LoadRepos ...
func (sources MultiSource) LoadRepos() ([]models.RepoState, error) {
	repos := []models.RepoState{}
	for _, source := range sources {
		repo, err := source.LoadRepos()
		if err != nil {
			return nil, fmt.Errorf("for the %s source: %v", source.Name(), err)
		}

		repos = append(repos, repo...)
	}

	return repos, nil
}
