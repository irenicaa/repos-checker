package gitlab

import (
	"fmt"
	"net/url"
	"strconv"
)

type project struct {
	RepoPath string `json:"path_with_namespace"`
}

// GetReposPage ...
func GetReposPage(owner string, isGroup bool, pageSize int, page int) (
	[]string,
	error,
) {
	parameters := url.Values{}
	parameters.Add("owned", "true")
	parameters.Add("order_by", "updated_at")
	parameters.Add("sort", "desc")
	parameters.Add("per_page", strconv.Itoa(pageSize))
	parameters.Add("page", strconv.Itoa(page))
	if isGroup {
		parameters.Add("include_subgroups", "true")
	}

	var entity string
	if isGroup {
		entity = "groups"
	} else {
		entity = "users"
	}

	var projects []project
	endpoint := fmt.Sprintf("/%s/%s/projects", entity, owner)
	if err := SendRequest(endpoint, parameters, &projects); err != nil {
		return nil, err
	}

	var reposPaths []string
	for _, project := range projects {
		reposPaths = append(reposPaths, project.RepoPath)
	}

	return reposPaths, nil
}
