package comparer

import "github.com/irenicaa/repos-checker/models"

// SourceDiff ...
type SourceDiff struct {
	NameOfLeft    string
	NameOfRight   string
	Equal         []models.RepoState
	Diff          []RepoDiff
	MissedInLeft  []models.RepoState
	MissedInRight []models.RepoState
}

// RepoDiff ...
type RepoDiff struct {
	Name              string
	LastCommitInLeft  string
	LastCommitInRight string
}

// CompareSources ...
func CompareSources(
	sources []models.SourceState,
	reference models.SourceState,
) []SourceDiff {
	sourceDiffs := []SourceDiff{}
	for _, source := range sources {
		sourceDiff := CompareRepos(source, reference)
		sourceDiffs = append(sourceDiffs, sourceDiff)
	}

	return sourceDiffs
}

// CompareRepos ...
func CompareRepos(
	left models.SourceState,
	right models.SourceState,
) SourceDiff {
	sourceDiff := SourceDiff{NameOfLeft: left.Name, NameOfRight: right.Name}

	rightRepos := models.NewRepoStateIndex(right.Repos)
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

	leftRepos := models.NewRepoStateIndex(left.Repos)
	for _, itemRight := range right.Repos {
		if _, found := leftRepos[itemRight.Name]; !found {
			sourceDiff.MissedInLeft = append(sourceDiff.MissedInLeft, itemRight)
		}
	}

	return sourceDiff
}
