package sourceutils

import (
	"testing"
	"testing/iotest"

	"github.com/irenicaa/repos-checker/v2/loader"
	"github.com/irenicaa/repos-checker/v2/models"
	"github.com/stretchr/testify/assert"
)

func TestLoadRepos(t *testing.T) {
	type args struct {
		sourceName    string
		getOnePage    GetOnePage
		getLastCommit GetLastCommit
		logger        loader.Logger
	}

	tests := []struct {
		name    string
		args    args
		want    []models.RepoState
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success without pages",
			args: args{
				sourceName:    "test",
				getOnePage:    func(page int) ([]string, error) { return nil, nil },
				getLastCommit: nil,
				logger:        &MockLogger{},
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "success with pages",
			args: args{
				sourceName: "test",
				getOnePage: func() GetOnePage {
					pages := [][]string{
						{"one", "two"},
						{"three", "four"},
						{"five", "six"},
						{"seven", "eight"},
						{"nine", "ten"},
					}
					return func(page int) ([]string, error) {
						pageIndex := page - 1
						if pageIndex > len(pages)-1 {
							return nil, nil
						}
						return pages[pageIndex], nil
					}
				}(),
				getLastCommit: func(repo string) (models.RepoState, error) {
					repoState := models.RepoState{
						Name:       repo,
						LastCommit: repo + "-last-commit",
					}
					return repoState, nil
				},
				logger: &MockLogger{},
			},
			want: []models.RepoState{
				{Name: "one", LastCommit: "one-last-commit"},
				{Name: "two", LastCommit: "two-last-commit"},
				{Name: "three", LastCommit: "three-last-commit"},
				{Name: "four", LastCommit: "four-last-commit"},
				{Name: "five", LastCommit: "five-last-commit"},
				{Name: "six", LastCommit: "six-last-commit"},
				{Name: "seven", LastCommit: "seven-last-commit"},
				{Name: "eight", LastCommit: "eight-last-commit"},
				{Name: "nine", LastCommit: "nine-last-commit"},
				{Name: "ten", LastCommit: "ten-last-commit"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "error with getting repos names",
			args: args{
				sourceName: "test",
				getOnePage: func() GetOnePage {
					pages := [][]string{
						{"one", "two"},
						{"three", "four"},
						{"five", "six"},
						{"seven", "eight"},
						{"nine", "ten"},
					}
					return func(page int) ([]string, error) {
						pageIndex := page - 1
						if pageIndex > len(pages)/2 {
							return nil, iotest.ErrTimeout
						}
						return pages[pageIndex], nil
					}
				}(),
				getLastCommit: nil,
				logger:        &MockLogger{},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "error without commits",
			args: args{
				sourceName: "test",
				getOnePage: func() GetOnePage {
					pages := [][]string{
						{"one", "two"},
						{"three", "four"},
						{"five", "six"},
						{"seven", "eight"},
						{"nine", "ten"},
					}
					return func(page int) ([]string, error) {
						pageIndex := page - 1
						if pageIndex > len(pages)-1 {
							return nil, nil
						}
						return pages[pageIndex], nil
					}
				}(),
				getLastCommit: func(repo string) (models.RepoState, error) {
					return models.RepoState{}, ErrNoCommits
				},
				logger: func() loader.Logger {
					repos := []string{
						"one", "two",
						"three", "four",
						"five", "six",
						"seven", "eight",
						"nine", "ten",
					}

					logger := &MockLogger{}
					for _, repo := range repos {
						arguments := []interface{}{repo, "test"}

						logger.InnerMock.
							On("Printf", "%s repo from the %s source has no commits", arguments).
							Return().
							Times(1)
					}

					return logger
				}(),
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "error with getting the last commit",
			args: args{
				sourceName: "test",
				getOnePage: func() GetOnePage {
					pages := [][]string{
						{"one", "two"},
						{"three", "four"},
						{"five", "six"},
						{"seven", "eight"},
						{"nine", "ten"},
					}
					return func(page int) ([]string, error) {
						pageIndex := page - 1
						if pageIndex > len(pages)-1 {
							return nil, nil
						}
						return pages[pageIndex], nil
					}
				}(),
				getLastCommit: func(repo string) (models.RepoState, error) {
					return models.RepoState{}, iotest.ErrTimeout
				},
				logger: &MockLogger{},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadRepos(
				tt.args.sourceName,
				tt.args.getOnePage,
				tt.args.getLastCommit,
				tt.args.logger,
			)

			tt.args.logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
