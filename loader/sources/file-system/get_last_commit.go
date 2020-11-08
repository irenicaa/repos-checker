package filesystem

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/irenicaa/repos-checker/models"
	systemutils "github.com/irenicaa/repos-checker/system-utils"
)

// GetLastCommit ...
func GetLastCommit(repoPath string) (models.RepoState, error) {
	commandOutput, err := systemutils.RunCommand(
		"git",
		[]string{"log", "--format=%H", "HEAD~.."},
		repoPath,
		nil,
	)
	if err != nil {
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
		LastCommit: strings.TrimSpace(string(commandOutput)),
	}

	return repoState, nil
}
