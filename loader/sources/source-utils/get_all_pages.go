package sourceutils

// GetOnePage ...
type GetOnePage func(page int) ([]string, error)

// GetAllPages ...
func GetAllPages(getOnePage GetOnePage) ([]string, error) {
	var allRepos []string
	for page := 1; ; page++ {
		repos, err := getOnePage(page)
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
