package models

import "reflect"

// SourceState ...
type SourceState struct {
	Name  string
	Repos []RepoState
}

// IsZero ...
func (sourceState SourceState) IsZero() bool {
	return reflect.DeepEqual(sourceState, SourceState{})
}

// FilterNonemptySources ...
func FilterNonemptySources(sourceStates []SourceState) []SourceState {
	filteredSourceState := []SourceState{}
	for _, sourceState := range sourceStates {
		if !sourceState.IsZero() {
			filteredSourceState = append(filteredSourceState, sourceState)
		}
	}

	return filteredSourceState
}
