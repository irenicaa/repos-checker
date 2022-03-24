package reposchecker

import (
	"github.com/irenicaa/repos-checker/v2/comparer"
	"github.com/irenicaa/repos-checker/v2/loader"
	"github.com/irenicaa/repos-checker/v2/models"
)

// CheckSources ...
func CheckSources(
	sources []loader.Source,
	referenceName string,
	logger loader.Logger,
) []comparer.SourceDiff {
	var reference models.SourceState
	var rest []models.SourceState
	sourceStates := loader.LoadSources(sources, logger)
	for _, sourceState := range sourceStates {
		duplicates := models.FindRepoStateDuplicates(sourceState.Repos)
		if len(duplicates) != 0 {
			logger.Printf(
				"repos from the %s source has duplicates: %v",
				sourceState.Name,
				duplicates,
			)
		}

		if sourceState.Name == referenceName {
			reference = sourceState
		} else {
			rest = append(rest, sourceState)
		}
	}
	if reference.IsZero() {
		logger.Printf("unable to load repos from the reference source")
		return []comparer.SourceDiff{}
	}

	return comparer.CompareSources(rest, reference)
}
