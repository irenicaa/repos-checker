package sourceutils

import (
	"errors"
	"fmt"

	"github.com/irenicaa/repos-checker/loader"
	"github.com/irenicaa/repos-checker/models"
)

// ErrNoCommits ...
var ErrNoCommits = errors.New("no commits")

// GetLastCommit ...
type GetLastCommit func(repo string) (models.RepoState, error)

// LoadRepos ...
func LoadRepos(
	getOnePage GetOnePage,
	getLastCommit GetLastCommit,
	logger loader.Logger,
) ([]models.RepoState, error) {
	repos, err := GetAllPages(getOnePage)
	if err != nil {
		return nil, fmt.Errorf("unable to get all repos names: %v", err)
	}

	var reposStates []models.RepoState
	for _, repo := range repos {
		repoState, err := getLastCommit(repo)
		switch err {
		case nil:
		case ErrNoCommits:
			logger.Printf("%s repo has no commits", repo)
		default:
			return nil, fmt.Errorf("unable to get the last commit: %v", err)
		}

		reposStates = append(reposStates, repoState)
	}

	return reposStates, nil
}
