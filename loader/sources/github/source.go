package github

import "github.com/irenicaa/repos-checker/models"

// Source ...
type Source struct {
	Owner string
}

// Name ...
func (source Source) Name() string {
	return "github"
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	const maxPageSize = 100
	repos, err := GetRepos(source.Owner, maxPageSize)
	if err != nil {
		return nil, err
	}

	var reposStates []models.RepoState
	for _, repo := range repos {
		repoState, err := GetLastCommit(source.Owner, repo)
		if err != nil {
			return nil, err
		}

		reposStates = append(reposStates, repoState)
	}

	return reposStates, nil
}
