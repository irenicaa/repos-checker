package loader

import (
	"errors"
	"fmt"
	"strings"
	"sync"

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
	waiter := sync.WaitGroup{}
	waiter.Add(len(sources))

	var mutex sync.Mutex
	repos := []models.RepoState{}
	errs := []error{}
	for _, source := range sources {
		go func(source Source) {
			defer waiter.Done()

			repo, err := source.LoadRepos()

			mutex.Lock()
			defer mutex.Unlock()

			repos = append(repos, repo...)
			if err != nil {
				err = fmt.Errorf("for the %s source: %v", source.Name(), err)
				errs = append(errs, err)
			}
		}(source)
	}

	waiter.Wait()
	if len(errs) != 0 {
		var errMessages []string
		for _, err := range errs {
			errMessages = append(errMessages, err.Error())
		}

		return nil, errors.New(strings.Join(errMessages, "; "))
	}

	return repos, nil
}
