package external

import (
	"encoding/json"
	"fmt"

	"github.com/irenicaa/repos-checker/models"
	systemutils "github.com/irenicaa/repos-checker/system-utils"
)

// Source ...
type Source struct {
	AdditionalName       string
	Command              string
	Arguments            []string
	WorkingDirectory     string
	EnvironmentVariables map[string]string
}

// Name ...
func (source Source) Name() string {
	return "external:" + source.AdditionalName
}

// LoadRepos ...
func (source Source) LoadRepos() ([]models.RepoState, error) {
	commandOutput, err := systemutils.RunCommand(
		source.Command,
		source.Arguments,
		source.WorkingDirectory,
		source.EnvironmentVariables,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"an error occurred while running the command: %v",
			err,
		)
	}

	var reposStates []models.RepoState
	if err := json.Unmarshal(commandOutput, &reposStates); err != nil {
		return nil, fmt.Errorf("unable to unmarshal a command response: %v", err)
	}

	return reposStates, nil
}
