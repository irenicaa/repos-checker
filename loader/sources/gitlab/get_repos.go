package gitlab

// GetRepos ...
func GetRepos(owner string, pageSize int) ([]string, error) {
	var allReposPaths []string
	for page := 1; ; page++ {
		reposPaths, err := GetReposPage(owner, pageSize, page)
		if err != nil {
			return nil, err
		}
		if len(reposPaths) == 0 {
			break
		}

		allReposPaths = append(allReposPaths, reposPaths...)
	}

	return allReposPaths, nil
}
