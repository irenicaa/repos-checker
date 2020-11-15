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
