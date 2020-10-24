package github

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/irenicaa/repos-checker/models"
)

type commit struct {
	SHA string `json:"sha"`
}

// GetLastCommit ...
func GetLastCommit(owner string, repo string) (models.RepoState, error) {
	parameters := url.Values{}
	parameters.Add("per_page", "1")

	var commits []commit
	endpoint := fmt.Sprintf("/repos/%s/%s/commits", owner, repo)
	if err := SendRequest(endpoint, parameters, &commits); err != nil {
		return models.RepoState{}, err
	}
	if len(commits) == 0 {
		return models.RepoState{}, errors.New("no commits")
	}

	repoState := models.RepoState{Name: repo, LastCommit: commits[0].SHA}
	return repoState, nil
}
