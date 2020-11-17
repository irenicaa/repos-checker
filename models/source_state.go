package models

import "reflect"

// SourceState ...
type SourceState struct {
	Name  string
	Repos []RepoState
}

// FilterNonemptySources ...
func FilterNonemptySources(sourceStates []SourceState) []SourceState {
	filteredSourceState := []SourceState{}
	for _, source := range sourceStates {
		if !reflect.DeepEqual(source, SourceState{}) {
			filteredSourceState = append(filteredSourceState, source)
		}
	}

	return filteredSourceState
}
