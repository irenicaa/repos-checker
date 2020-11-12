package filesystem

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	sourceutils "github.com/irenicaa/repos-checker/loader/sources/source-utils"
	"github.com/irenicaa/repos-checker/models"
	systemutils "github.com/irenicaa/repos-checker/system-utils"
)

// CheckCommitCount ...
func CheckCommitCount(repoPath string) error {
	statusOutput, err := systemutils.RunCommand(
		"git",
		[]string{"status"},
		repoPath,
		map[string]string{"LC_ALL": "en_US"},
	)
	if err != nil {
		return fmt.Errorf("unable to get a git status: %v", err)
	}

	statusOutput = bytes.ToLower(statusOutput)
	if bytes.Index(statusOutput, []byte("no commits yet")) != -1 {
		return sourceutils.ErrNoCommits
	}

	return nil
}

// GetLastCommit ...
func GetLastCommit(repoPath string) (models.RepoState, error) {
	if err := CheckCommitCount(repoPath); err != nil {
		return models.RepoState{}, err
	}

	logOutput, err := systemutils.RunCommand(
		"git",
		[]string{"log", "--format=%H", "HEAD~.."},
		repoPath,
		nil,
	)
	if err != nil {
		return models.RepoState{}, fmt.Errorf("unable to get a git log: %v", err)
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
		LastCommit: strings.TrimSpace(string(logOutput)),
	}

	return repoState, nil
}
