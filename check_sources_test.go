package reposchecker

import (
	"reflect"
	"testing"

	"github.com/irenicaa/repos-checker/v2/comparer"
	"github.com/irenicaa/repos-checker/v2/loader"
	"github.com/irenicaa/repos-checker/v2/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckSources(t *testing.T) {
	type args struct {
		sources       []loader.Source
		referenceName string
		logger        loader.Logger
	}

	tests := []struct {
		name string
		args args
		want []comparer.SourceDiff
	}{
		{
			name: "empty",
			args: args{
				sources:       []loader.Source{},
				referenceName: "source-three",
				logger: func() loader.Logger {
					arguments := []interface{}(nil)

					logger := &MockLogger{}
					logger.InnerMock.
						On("Printf", "unable to load repos from the reference source", arguments).
						Return().
						Times(1)

					return logger
				}(),
			},
			want: []comparer.SourceDiff{},
		},
		{
			name: "without a reference",
			args: args{
				sources: []loader.Source{
					func() loader.Source {
						repos := []models.RepoState{
							{Name: "one", LastCommit: "100"},
							{Name: "two", LastCommit: "200"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-one").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
					func() loader.Source {
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
				referenceName: "source-three",
				logger: func() loader.Logger {
					arguments := []interface{}(nil)

					logger := &MockLogger{}
					logger.InnerMock.
						On("Printf", "unable to load repos from the reference source", arguments).
						Return().
						Times(1)

					return logger
				}(),
			},
			want: []comparer.SourceDiff{},
		},
		{
			name: "with a reference",
			args: args{
				sources: []loader.Source{
					func() loader.Source {
						repos := []models.RepoState{
							{Name: "one", LastCommit: "100"},
							{Name: "two", LastCommit: "200"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-one").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
					func() loader.Source {
						repos := []models.RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-two").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
					func() loader.Source {
						repos := []models.RepoState{
							{Name: "five", LastCommit: "500"},
							{Name: "six", LastCommit: "600"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-three").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
				},
				referenceName: "source-three",
				logger:        &MockLogger{},
			},
			want: []comparer.SourceDiff{
				{
					NameOfLeft:  "source-one",
					NameOfRight: "source-three",
					MissedInLeft: []models.RepoState{
						{Name: "five", LastCommit: "500"},
						{Name: "six", LastCommit: "600"},
					},
					MissedInRight: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				{
					NameOfLeft:  "source-two",
					NameOfRight: "source-three",
					MissedInLeft: []models.RepoState{
						{Name: "five", LastCommit: "500"},
						{Name: "six", LastCommit: "600"},
					},
					MissedInRight: []models.RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
			},
		},
		{
			name: "with duplicated repos",
			args: args{
				sources: []loader.Source{
					func() loader.Source {
						repos := []models.RepoState{
							{Name: "one", LastCommit: "100"},
							{Name: "two", LastCommit: "200"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-one").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
					func() loader.Source {
						repos := []models.RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "three", LastCommit: "350"},
							{Name: "four", LastCommit: "400"},
							{Name: "four", LastCommit: "450"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-two").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
					func() loader.Source {
						repos := []models.RepoState{
							{Name: "five", LastCommit: "500"},
							{Name: "six", LastCommit: "600"},
						}

						source := &MockSource{}
						source.InnerMock.On("Name").Return("source-three").Times(1)
						source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

						return source
					}(),
				},
				referenceName: "source-three",
				logger: func() loader.Logger {
					argumentsMatcher := mock.MatchedBy(func(arguments []interface{}) bool {
						argumentsOne := []interface{}{"source-two", []string{"three", "four"}}
						argumentsTwo := []interface{}{"source-two", []string{"four", "three"}}

						return reflect.DeepEqual(arguments, argumentsOne) ||
							reflect.DeepEqual(arguments, argumentsTwo)
					})

					logger := &MockLogger{}
					logger.InnerMock.
						On(
							"Printf",
							"repos from the %s source has duplicates: %v",
							argumentsMatcher,
						).
						Return().
						Times(1)

					return logger
				}(),
			},
			want: []comparer.SourceDiff{
				{
					NameOfLeft:  "source-one",
					NameOfRight: "source-three",
					MissedInLeft: []models.RepoState{
						{Name: "five", LastCommit: "500"},
						{Name: "six", LastCommit: "600"},
					},
					MissedInRight: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				{
					NameOfLeft:  "source-two",
					NameOfRight: "source-three",
					MissedInLeft: []models.RepoState{
						{Name: "five", LastCommit: "500"},
						{Name: "six", LastCommit: "600"},
					},
					MissedInRight: []models.RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "three", LastCommit: "350"},
						{Name: "four", LastCommit: "400"},
						{Name: "four", LastCommit: "450"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckSources(tt.args.sources, tt.args.referenceName, tt.args.logger)

			for _, source := range tt.args.sources {
				source.(*MockSource).InnerMock.AssertExpectations(t)
			}
			tt.args.logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.want, got)
		})
	}
}
