package gitlab

import (
	"fmt"
	"net/url"
	"path"

	sourceutils "github.com/irenicaa/repos-checker/loader/sources/source-utils"
	"github.com/irenicaa/repos-checker/models"
)

type commit struct {
	SHA string `json:"id"`
}

// GetLastCommit ...
func GetLastCommit(repoPath string) (models.RepoState, error) {
	parameters := url.Values{}
	parameters.Add("per_page", "1")

	var commits []commit
	escapedRepoPath := url.PathEscape(repoPath)
	endpoint := fmt.Sprintf("/projects/%s/repository/commits", escapedRepoPath)
	if err := SendRequest(endpoint, parameters, &commits); err != nil {
		return models.RepoState{}, err
	}
	if len(commits) == 0 {
		return models.RepoState{}, sourceutils.ErrNoCommits
	}

	_, repo := path.Split(repoPath)
	repoState := models.RepoState{Name: repo, LastCommit: commits[0].SHA}
	return repoState, nil
}
