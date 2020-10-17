package loader

import (
	"testing"
	"testing/iotest"

	"github.com/irenicaa/repos-checker/models"
	"github.com/stretchr/testify/assert"
)

func TestLoadSources(t *testing.T) {
	type args struct {
		sources []Source
		logger  Logger
	}

	tests := []struct {
		name string
		args args
		want []models.SourceState
	}{
		{
			name: "empty",
			args: args{sources: []Source{}, logger: &MockLogger{}},
			want: []models.SourceState{},
		},
		{
			name: "nonempty",
			args: args{
				sources: []Source{
					func() Source {
						repos := []models.RepoState{
							{Name: "one", LastCommit: "100"},
							{Name: "two", LastCommit: "200"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-one").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
					func() Source {
						repos := []models.RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-two").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
				},
				logger: &MockLogger{},
			},
			want: []models.SourceState{
				models.SourceState{
					Name: "source-one",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				models.SourceState{
					Name: "source-two",
					Repos: []models.RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
			},
		},
		{
			name: "few failed sources",
			args: args{
				sources: []Source{
					func() Source {
						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-one").Times(1)
						source.InnerMock.
							On("LoadRepos").
							Return(([]models.RepoState)(nil), iotest.ErrTimeout).
							Times(1)

						return source
					}(),
					func() Source {
						repos := []models.RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-two").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
				},
				logger: func() Logger {
					arguments := []interface{}{"source-one", iotest.ErrTimeout}

					logger := &MockLogger{}
					logger.InnerMock.
						On("Printf", "unable to load repos from the %s source: %s", arguments).
						Return().
						Times(1)

					return logger
				}(),
			},
			want: []models.SourceState{
				models.SourceState{
					Name: "source-two",
					Repos: []models.RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
			},
		},
		{
			name: "all failed sources",
			args: args{
				sources: []Source{
					func() Source {
						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-one").Times(1)
						source.InnerMock.
							On("LoadRepos").
							Return(([]models.RepoState)(nil), iotest.ErrTimeout).
							Times(1)

						return source
					}(),
					func() Source {
						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-two").Times(1)
						source.InnerMock.On("LoadRepos").
							Return(([]models.RepoState)(nil), iotest.ErrTimeout).
							Times(1)

						return source
					}(),
				},
				logger: func() Logger {
					argumentsOne := []interface{}{"source-one", iotest.ErrTimeout}
					argumentsTwo := []interface{}{"source-two", iotest.ErrTimeout}

					logger := &MockLogger{}
					logger.InnerMock.
						On("Printf", "unable to load repos from the %s source: %s", argumentsOne).
						Return().
						Times(1)
					logger.InnerMock.
						On("Printf", "unable to load repos from the %s source: %s", argumentsTwo).
						Return().
						Times(1)

					return logger
				}(),
			},
			want: []models.SourceState{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadSources(tt.args.sources, tt.args.logger)

			for _, source := range tt.args.sources {
				source.(*MockSource).InnerMock.AssertExpectations(t)
			}
			tt.args.logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.want, got)
		})
	}
}
