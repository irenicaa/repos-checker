package systemutils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// RunCommand ...
func RunCommand(
	name string,
	arguments []string,
	workingDirectory string,
	environmentVariables map[string]string,
) ([]byte, error) {
	command := exec.Command(name, arguments...)
	command.Dir = workingDirectory

	for key, value := range environmentVariables {
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
		return nil, err
	}

	return stdoutBuffer.Bytes(), nil
}
