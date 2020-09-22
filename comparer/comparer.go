package comparer

import "github.com/irenicaa/repos-checker/models"

// RepoDiff ...
type RepoDiff struct {
	Name              string
	LastCommitInLeft  string
	LastCommitInRight string
}

// SourceDiff ...
type SourceDiff struct {
	Equal         []models.RepoState
	Diff          []RepoDiff
	MissedInLeft  []models.RepoState
	MissedInRight []models.RepoState
}

// CompareRepos ...
func CompareRepos(
	left models.SourceState,
	right models.SourceState,
) SourceDiff {
	sourceDiff := SourceDiff{}
	for _, itemLeft := range left.Repos {
		for _, itemRight := range right.Repos {
			if itemLeft.Name == itemRight.Name {
				if itemLeft.LastCommit == itemRight.LastCommit {
					sourceDiff.Equal = append(sourceDiff.Equal, itemLeft)
				} else {
					sourceDiff.Diff = append(sourceDiff.Diff, RepoDiff{
						Name:              itemLeft.Name,
						LastCommitInLeft:  itemLeft.LastCommit,
						LastCommitInRight: itemRight.LastCommit,
					})
				}
			} else {
				sourceDiff.MissedInRight = append(sourceDiff.MissedInRight, itemLeft)
			}
		}
	}
	for _, itemRight := range right.Repos {
		found := false
		for _, itemLeft := range left.Repos {
			if itemRight.Name == itemLeft.Name {
				found = true
				break
			}
		}
		if !found {
			sourceDiff.MissedInLeft = append(sourceDiff.MissedInRight, itemRight)
		}
	}

	return sourceDiff
}
