package bitbucket

import "github.com/irenicaa/repos-checker/models"

// Source ...
type Source struct {
	Workspace string
}

// Name ...
func (source Source) Name() string {
	return "bitbucket"
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	const maxPageSize = 100
	repos, err := GetRepos(source.Workspace, maxPageSize)
	if err != nil {
		return nil, err
	}

	var reposStates []models.RepoState
	for _, repo := range repos {
		repoState, err := GetLastCommit(source.Workspace, repo)
		if err != nil {
			return nil, err
		}

		reposStates = append(reposStates, repoState)
	}

	return reposStates, nil
}
