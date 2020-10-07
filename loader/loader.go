package loader

import "github.com/irenicaa/repos-checker/models"

// Source ...
type Source interface {
	Name() string
	LoadRepos() []models.RepoState
}

// LoadSources ...
func LoadSources(sources []Source) []models.SourceState {
	sourceStates := make([]models.SourceState, len(sources))
	for index, source := range sources {
		sourceStates[index] = models.SourceState{
			Name:  source.Name(),
			Repos: source.LoadRepos(),
		}
	}

	return sourceStates
}
