package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSourceState_IsZero(t *testing.T) {
	tests := []struct {
		name        string
		sourceState SourceState
		want        bool
	}{
		{
			name: "with full data",
			sourceState: SourceState{
				Name: "source-one",
				Repos: []RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
			},
			want: false,
		},
		{
			name: "without a name",
			sourceState: SourceState{
				Name: "",
				Repos: []RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
			},
			want: false,
		},
		{
			name: "without repos",
			sourceState: SourceState{
				Name:  "source-one",
				Repos: nil,
			},
			want: false,
		},
		{
			name: "with an empty source state",
			sourceState: SourceState{
				Name:  "",
				Repos: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sourceState.IsZero()

			assert.Equal(t, tt.want, got)
		})
	}
}

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
					{
						Name: "source-one",
						Repos: []RepoState{
							{Name: "one", LastCommit: "100"},
							{Name: "two", LastCommit: "200"},
						},
					},
					{
						Name: "source-two",
						Repos: []RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						},
					},
				},
			},
			want: []SourceState{
				{
					Name: "source-one",
					Repos: []RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				{
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
					{
						Name: "",
						Repos: []RepoState{
							{Name: "one", LastCommit: "100"},
							{Name: "two", LastCommit: "200"},
						},
					},
					{
						Name: "source-two",
						Repos: []RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						},
					},
				},
			},
			want: []SourceState{
				{
					Name: "",
					Repos: []RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				{
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
					{
						Name:  "source-one",
						Repos: nil,
					},
					{
						Name: "source-two",
						Repos: []RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						},
					},
				},
			},
			want: []SourceState{
				{
					Name:  "source-one",
					Repos: nil,
				},
				{
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
					{
						Name:  "",
						Repos: nil,
					},
					{
						Name: "source-two",
						Repos: []RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						},
					},
				},
			},
			want: []SourceState{
				{
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
					{
						Name:  "",
						Repos: nil,
					},
					{
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
