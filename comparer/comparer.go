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
	NameOfLeft    string
	NameOfRight   string
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
	leftRepos := makeRepoIndex(left)

	rightRepos := makeRepoIndex(right)

	sourceDiff := SourceDiff{NameOfLeft: left.Name, NameOfRight: right.Name}
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

func makeRepoIndex(repos models.SourceState) map[string]models.RepoState {
	repoIndex := map[string]models.RepoState{}
	for _, item := range repos.Repos {
		repoIndex[item.Name] = item
	}

	return repoIndex
}
