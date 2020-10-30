package bitbucket

import (
	"fmt"
	"net/url"
	"strconv"
)

type reposPage struct {
	Values []repo
}

type repo struct {
	Name string `json:"name"`
}

// GetReposPage ...
func GetReposPage(workspace string, pageSize int, page int) ([]string, error) {
	parameters := url.Values{}
	parameters.Add("sort", "-updated_on")
	parameters.Add("pagelen", strconv.Itoa(pageSize))
	parameters.Add("page", strconv.Itoa(page))

	var repos reposPage
	endpoint := fmt.Sprintf("/repositories/%s", workspace)
	if err := SendRequest(endpoint, parameters, &repos); err != nil {
		return nil, err
	}

	var reposNames []string
	for _, repo := range repos.Values {
		reposNames = append(reposNames, repo.Name)
	}

	return reposNames, nil
}
