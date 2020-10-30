package bitbucket

// GetRepos ...
func GetRepos(workspace string, pageSize int) ([]string, error) {
	var allRepos []string
	for page := 1; ; page++ {
		repos, err := GetReposPage(workspace, pageSize, page)
		if err != nil {
			return nil, err
		}
		if len(repos) == 0 {
			break
		}

		allRepos = append(allRepos, repos...)
	}

	return allRepos, nil
}
