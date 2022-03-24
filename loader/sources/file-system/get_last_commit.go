package filesystem

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	sourceutils "github.com/irenicaa/repos-checker/v2/loader/sources/source-utils"
	"github.com/irenicaa/repos-checker/v2/models"
	systemutils "github.com/irenicaa/repos-checker/v2/system-utils"
)

// GetRepoName ...
func GetRepoName(repoPath string) (string, error) {
	absoluteRepoPath, err := filepath.Abs(repoPath)
	if err != nil {
		return "", fmt.Errorf("unable to get an absolute repo path: %v", err)
	}

	_, repoName := filepath.Split(absoluteRepoPath)
	return repoName, nil
}

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

// GetLastCommitSHA ...
func GetLastCommitSHA(repoPath string) (string, error) {
	logOutput, err := systemutils.RunCommand(
		"git",
		[]string{"rev-parse", "HEAD"},
		repoPath,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("unable to parse a HEAD git reference: %v", err)
	}

	return strings.TrimSpace(string(logOutput)), nil
}

// GetLastCommit ...
func GetLastCommit(repoPath string) (models.RepoState, error) {
	repoName, err := GetRepoName(repoPath)
	if err != nil {
		return models.RepoState{}, err
	}

	if err = CheckCommitCount(repoPath); err != nil {
		return models.RepoState{}, err
	}

	lastCommitSHA, err := GetLastCommitSHA(repoPath)
	if err != nil {
		return models.RepoState{}, err
	}

	repoState := models.RepoState{Name: repoName, LastCommit: lastCommitSHA}
	return repoState, nil
}
