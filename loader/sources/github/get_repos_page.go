package github

import (
	"fmt"
	"net/url"
	"strconv"
)

type repo struct {
	Name string `json:"name"`
}

// GetReposPage ...
func GetReposPage(owner string, pageSize int, page int) ([]string, error) {
	parameters := url.Values{}
	parameters.Add("type", "owner")
	parameters.Add("sort", "pushed")
	parameters.Add("per_page", strconv.Itoa(pageSize))
	parameters.Add("page", strconv.Itoa(page))

	var repos []repo
	endpoint := fmt.Sprintf("/users/%s/repos", owner)
	if err := LoadData(endpoint, parameters, &repos); err != nil {
		return nil, err
	}

	var reposNames []string
	for _, repo := range repos {
		reposNames = append(reposNames, repo.Name)
	}

	return reposNames, nil
}
