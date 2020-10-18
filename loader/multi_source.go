package loader

import "strings"

// MultiSource ...
type MultiSource []Source

// Name ...
func (sources MultiSource) Name() string {
	names := []string{}
	for _, source := range sources {
		name := source.Name()
		names = append(names, name)
	}

	return strings.Join(names, "|")
}
