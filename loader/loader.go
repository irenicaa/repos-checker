package loader

import "github.com/irenicaa/repos-checker/models"

// Source ...
type Source interface {
	Name() string
	LoadRepos() []models.RepoState
}

// LoadSources ...
func LoadSources(sources []Source) []models.SourceState {
	sourceStates := []models.SourceState{}
	for _, source := range sources {
		sourceStates = append(sourceStates, models.SourceState{
			Name:  source.Name(),
			Repos: source.LoadRepos(),
		})
	}

	return sourceStates
}
