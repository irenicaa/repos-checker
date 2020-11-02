package bitbucket

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/irenicaa/repos-checker/models"
)

type commitsPage struct {
	Values []commit
}

type commit struct {
	SHA string `json:"hash"`
}

var errNoCommits = errors.New("no commits")

// GetLastCommit ...
func GetLastCommit(workspace string, repo string) (models.RepoState, error) {
	parameters := url.Values{}
	parameters.Add("pagelen", "1")

	var commits commitsPage
	endpoint := fmt.Sprintf("/repositories/%s/%s/commits", workspace, repo)
	if err := SendRequest(endpoint, parameters, &commits); err != nil {
		return models.RepoState{}, err
	}
	if len(commits.Values) == 0 {
		return models.RepoState{}, errNoCommits
	}

	repoState := models.RepoState{Name: repo, LastCommit: commits.Values[0].SHA}
	return repoState, nil
}
