package models

// RepoState ...
type RepoState struct {
	Name       string
	LastCommit string
}

// SourceState ...
type SourceState struct {
	Name  string
	Repos []RepoState
}
