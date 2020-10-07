package loader

import (
	"sync"

	"github.com/irenicaa/repos-checker/models"
)

// Source ...
type Source interface {
	Name() string
	LoadRepos() []models.RepoState
}

// LoadSources ...
func LoadSources(sources []Source) []models.SourceState {
	waiter := sync.WaitGroup{}
	waiter.Add(len(sources))

	sourceStates := make([]models.SourceState, len(sources))
	for index, source := range sources {
		go func(index int, source Source) {
			defer waiter.Done()

			sourceStates[index] = models.SourceState{
				Name:  source.Name(),
				Repos: source.LoadRepos(),
			}
		}(index, source)
	}

	waiter.Wait()
	return sourceStates
}
