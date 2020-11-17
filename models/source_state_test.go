package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterNonemptySources(t *testing.T) {
	type args struct {
		sourceStates []SourceState
	}

	tests := []struct {
		name string
		args args
		want []SourceState
	}{
		{
			name: "empty",
			args: args{sourceStates: []SourceState{}},
			want: []SourceState{},
		},
		{
			name: "with full data",
			args: args{
				sourceStates: []SourceState{
					SourceState{
						Name: "source-one",
						Repos: []RepoState{
							{Name: "one", LastCommit: "100"},
							{Name: "two", LastCommit: "200"},
						},
					},
					SourceState{
						Name: "source-two",
						Repos: []RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						},
					},
				},
			},
			want: []SourceState{
				SourceState{
					Name: "source-one",
					Repos: []RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				SourceState{
					Name: "source-two",
					Repos: []RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
			},
		},
		{
			name: "without a name",
			args: args{
				sourceStates: []SourceState{
					SourceState{
						Name: "",
						Repos: []RepoState{
							{Name: "one", LastCommit: "100"},
							{Name: "two", LastCommit: "200"},
						},
					},
					SourceState{
						Name: "source-two",
						Repos: []RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						},
					},
				},
			},
			want: []SourceState{
				SourceState{
					Name: "",
					Repos: []RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				SourceState{
					Name: "source-two",
					Repos: []RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
			},
		},
		{
			name: "without repos",
			args: args{
				sourceStates: []SourceState{
					SourceState{
						Name:  "source-one",
						Repos: nil,
					},
					SourceState{
						Name: "source-two",
						Repos: []RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						},
					},
				},
			},
			want: []SourceState{
				SourceState{
					Name:  "source-one",
					Repos: nil,
				},
				SourceState{
					Name: "source-two",
					Repos: []RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
			},
		},
		{
			name: "with an empty source state",
			args: args{
				sourceStates: []SourceState{
					SourceState{
						Name:  "",
						Repos: nil,
					},
					SourceState{
						Name: "source-two",
						Repos: []RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						},
					},
				},
			},
			want: []SourceState{
				SourceState{
					Name: "source-two",
					Repos: []RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
			},
		},
		{
			name: "without a nonempty source state",
			args: args{
				sourceStates: []SourceState{
					SourceState{
						Name:  "",
						Repos: nil,
					},
					SourceState{
						Name:  "",
						Repos: nil,
					},
				},
			},
			want: []SourceState{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterNonemptySources(tt.args.sourceStates)

			assert.Equal(t, tt.want, got)
		})
	}
}
