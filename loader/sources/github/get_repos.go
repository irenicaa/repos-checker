package github

// GetRepos ...
func GetRepos(owner string, pageSize int) ([]string, error) {
	var allRepos []string
	for page := 1; ; page++ {
		repos, err := GetReposPage(owner, pageSize, page)
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
