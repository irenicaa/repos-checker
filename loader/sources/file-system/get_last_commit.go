package filesystem

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/irenicaa/repos-checker/models"
)

// GetLastCommit ...
func GetLastCommit(repoPath string) (models.RepoState, error) {
	command := exec.Command("git", "log", "--format=%H", "HEAD~..")
	command.Dir = repoPath

	var stdoutBuffer bytes.Buffer
	command.Stdout = &stdoutBuffer

	var stderrBuffer bytes.Buffer
	command.Stderr = &stderrBuffer

	if err := command.Run(); err != nil {
		if errMessage := stderrBuffer.String(); errMessage != "" {
			err = fmt.Errorf("%v: %q", err, stderrBuffer.String())
		}
		return models.RepoState{}, fmt.Errorf(
			"an error occurred while running the command: %v",
			err,
		)
	}

	absoluteRepoPath, err := filepath.Abs(repoPath)
	if err != nil {
		return models.RepoState{}, fmt.Errorf(
			"unable to get an absolute repo path: %v",
			err,
		)
	}

	_, repo := filepath.Split(absoluteRepoPath)
	repoState := models.RepoState{
		Name:       repo,
		LastCommit: strings.TrimSpace(stdoutBuffer.String()),
	}

	return repoState, nil
}
