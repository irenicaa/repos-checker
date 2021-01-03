package external

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/irenicaa/repos-checker/models"
	systemutils "github.com/irenicaa/repos-checker/system-utils"
)

// Source ...
type Source struct {
	Command              string
	Arguments            []string
	WorkingDirectory     string
	EnvironmentVariables map[string]string
}

// Name ...
func (source Source) Name() string {
	environmentVariablesPairs :=
		systemutils.PrepareEnvironmentVariables(source.EnvironmentVariables)
	environmentVariables := strings.Join(environmentVariablesPairs, " ")

	arguments := strings.Join(source.Arguments, " ")

	return fmt.Sprintf(
		"external:%s:%s %s %s",
		source.WorkingDirectory,
		environmentVariables,
		source.Command,
		arguments,
	)
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
		return nil, fmt.Errorf("error occurred while running the command: %v", err)
	}

	var reposStates []models.RepoState
	if err := json.Unmarshal(commandOutput, &reposStates); err != nil {
		return nil, fmt.Errorf("unable to unmarshal a command response: %v", err)
	}

	return reposStates, nil
}
