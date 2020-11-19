package models

// RepoState ...
type RepoState struct {
	Name       string
	LastCommit string
}

// RepoStateIndex ...
type RepoStateIndex map[string]RepoState

// NewRepoStateIndex ...
func NewRepoStateIndex(repos []RepoState) RepoStateIndex {
	repoIndex := RepoStateIndex{}
	for _, item := range repos {
		repoIndex[item.Name] = item
	}

	return repoIndex
}

// FindRepoStateDuplicates ...
func FindRepoStateDuplicates(reposStates []RepoState) []string {
	reposStatesCounters := map[string]int{}
	for _, repoState := range reposStates {
		reposStatesCounters[repoState.Name]++
	}

	var duplicates []string
	for name, counter := range reposStatesCounters {
		if counter > 1 {
			duplicates = append(duplicates, name)
		}
	}

	return duplicates
}
