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
	leftRepos := map[string]models.RepoState{}
	for _, item := range left.Repos {
		leftRepos[item.Name] = item
	}

	rightRepos := map[string]models.RepoState{}
	for _, item := range right.Repos {
		rightRepos[item.Name] = item
	}

	sourceDiff := SourceDiff{}
	for _, itemLeft := range left.Repos {
		foundItem, found := rightRepos[itemLeft.Name]
		if !found {
			sourceDiff.MissedInRight = append(sourceDiff.MissedInRight, itemLeft)
			continue
		}

		if itemLeft.LastCommit == foundItem.LastCommit {
			sourceDiff.Equal = append(sourceDiff.Equal, itemLeft)
		} else {
			sourceDiff.Diff = append(sourceDiff.Diff, RepoDiff{
				Name:              itemLeft.Name,
				LastCommitInLeft:  itemLeft.LastCommit,
				LastCommitInRight: foundItem.LastCommit,
			})
		}
	}
	for _, itemRight := range right.Repos {
		_, found := leftRepos[itemRight.Name]
		if !found {
			sourceDiff.MissedInLeft = append(sourceDiff.MissedInRight, itemRight)
		}
	}

	return sourceDiff
}
