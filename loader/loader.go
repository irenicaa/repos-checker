package loader

import (
	"sync"

	"github.com/irenicaa/repos-checker/models"
)

// Source ...
type Source interface {
	Name() string
	LoadRepos() ([]models.RepoState, error)
}

// Logger ...
type Logger interface {
	Printf(format string, arguments ...interface{})
}

// LoadSources ...
func LoadSources(sources []Source, logger Logger) []models.SourceState {
	waiter := sync.WaitGroup{}
	waiter.Add(len(sources))

	sourceStates := make([]models.SourceState, len(sources))
	for index, source := range sources {
		go func(index int, source Source) {
			defer waiter.Done()

			name := source.Name()
			repos, err := source.LoadRepos()
			if err != nil {
				logger.Printf("unable to load repos from the %s source: %s", name, err)
				return
			}

			sourceStates[index] = models.SourceState{
				Name:  name,
				Repos: repos,
			}
		}(index, source)
	}

	waiter.Wait()
	return sourceStates
}
