package models

// SourceState ...
type SourceState struct {
	Name  string
	Repos []RepoState
}

// RepoState ...
type RepoState struct {
	Name       string
	LastCommit string
}
