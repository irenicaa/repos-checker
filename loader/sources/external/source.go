package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/irenicaa/repos-checker/models"
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
	command := exec.Command(source.Command, source.Arguments...)
	command.Dir = source.WorkingDirectory

	for key, value := range source.EnvironmentVariables {
		entry := key + "=" + value
		command.Env = append(command.Env, entry)
	}

	var stdoutBuffer bytes.Buffer
	command.Stdout = &stdoutBuffer

	var stderrBuffer bytes.Buffer
	command.Stderr = &stderrBuffer

	if err := command.Run(); err != nil {
		if errMessage := stderrBuffer.String(); errMessage != "" {
			err = fmt.Errorf("%v: %q", err, stderrBuffer.String())
		}
		return nil, fmt.Errorf("an error occurred while running the command: %v", err)
	}

	var reposStates []models.RepoState
	if err := json.Unmarshal(stdoutBuffer.Bytes(), &reposStates); err != nil {
		return nil, fmt.Errorf("unable to unmarshal a command response: %v", err)
	}

	return reposStates, nil
}
