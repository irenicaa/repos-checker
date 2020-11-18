package reposchecker

import (
	"reflect"

	"github.com/irenicaa/repos-checker/comparer"
	"github.com/irenicaa/repos-checker/loader"
	"github.com/irenicaa/repos-checker/models"
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
	for _, sourseState := range sourceStates {
		if sourseState.Name == referenceName {
			reference = sourseState
		} else {
			rest = append(rest, sourseState)
		}
	}
	if reflect.DeepEqual(reference, models.SourceState{}) {
		logger.Printf("no repos for the reference source")
		return []comparer.SourceDiff{}
	}

	return comparer.CompareSources(rest, reference)
}
