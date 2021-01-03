package systemutils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// PrepareEnvironmentVariables ...
func PrepareEnvironmentVariables(
	environmentVariables map[string]string,
) []string {
	var environmentVariablesPairs []string
	for name, value := range environmentVariables {
		pair := name + "=" + value
		environmentVariablesPairs = append(environmentVariablesPairs, pair)
	}

	return environmentVariablesPairs
}

// RunCommand ...
func RunCommand(
	name string,
	arguments []string,
	workingDirectory string,
	environmentVariables map[string]string,
) ([]byte, error) {
	command := exec.Command(name, arguments...)
	command.Dir = workingDirectory
	command.Env = PrepareEnvironmentVariables(environmentVariables)

	var stdoutBuffer bytes.Buffer
	command.Stdout = &stdoutBuffer

	var stderrBuffer bytes.Buffer
	command.Stderr = &stderrBuffer

	if err := command.Run(); err != nil {
		if errMessage := stderrBuffer.String(); errMessage != "" {
			err = fmt.Errorf("%v: %q", err, stderrBuffer.String())
		}
		return nil, err
	}

	return stdoutBuffer.Bytes(), nil
}
