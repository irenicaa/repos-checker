package loader

import (
	"sync"

	"github.com/irenicaa/repos-checker/v2/models"
)

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

			sourceState, err := LoadSource(source)
			if err != nil {
				logger.Printf(
					"unable to load repos from the %s source: %v",
					sourceState.Name,
					err,
				)

				return
			}

			sourceStates[index] = sourceState
		}(index, source)
	}

	waiter.Wait()

	return models.FilterNonemptySources(sourceStates)
}

// LoadSource ...
func LoadSource(source Source) (models.SourceState, error) {
	name := source.Name()
	repos, err := source.LoadRepos()
	if err != nil {
		return models.SourceState{Name: name}, err
	}

	sourceState := models.SourceState{Name: name, Repos: repos}
	return sourceState, nil
}
