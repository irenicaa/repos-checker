package comparer

import (
	"testing"

	"github.com/irenicaa/repos-checker/models"
	"github.com/stretchr/testify/assert"
)

func TestCompareSources(t *testing.T) {
	type args struct {
		sources   []models.SourceState
		reference models.SourceState
	}

	tests := []struct {
		name string
		args args
		want []SourceDiff
	}{
		{
			name: "empty",
			args: args{
				sources: []models.SourceState{},
				reference: models.SourceState{
					Name: "reference",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
			},
			want: []SourceDiff{},
		},
		{
			name: "nonempty",
			args: args{
				sources: []models.SourceState{
					{
						Name: "source-one",
						Repos: []models.RepoState{
							{Name: "one", LastCommit: "100"},
							{Name: "two", LastCommit: "250"},
						},
					},
					{
						Name: "source-two",
						Repos: []models.RepoState{
							{Name: "three", LastCommit: "300"},
							{Name: "four", LastCommit: "400"},
						},
					},
				},
				reference: models.SourceState{
					Name: "reference",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
			},
			want: []SourceDiff{
				{
					NameOfLeft:  "source-one",
					NameOfRight: "reference",
					Equal: []models.RepoState{
						{Name: "one", LastCommit: "100"},
					},
					Diff: []RepoDiff{
						{Name: "two", LastCommitInLeft: "250", LastCommitInRight: "200"},
					},
				},
				{
					NameOfLeft:  "source-two",
					NameOfRight: "reference",
					MissedInLeft: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
					MissedInRight: []models.RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompareSources(tt.args.sources, tt.args.reference)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCompareRepos(t *testing.T) {
	type args struct {
		left  models.SourceState
		right models.SourceState
	}

	tests := []struct {
		name string
		args args
		want SourceDiff
	}{
		{
			name: "both are empty",
			args: args{
				left:  models.SourceState{Name: "left"},
				right: models.SourceState{Name: "right"},
			},
			want: SourceDiff{NameOfLeft: "left", NameOfRight: "right"},
		},
		{
			name: "left is empty",
			args: args{
				left: models.SourceState{Name: "left"},
				right: models.SourceState{
					Name: "right",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
			},
			want: SourceDiff{
				NameOfLeft:  "left",
				NameOfRight: "right",
				MissedInLeft: []models.RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
			},
		},
		{
			name: "right is empty",
			args: args{
				left: models.SourceState{
					Name: "left",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				right: models.SourceState{Name: "right"},
			},
			want: SourceDiff{
				NameOfLeft:  "left",
				NameOfRight: "right",
				MissedInRight: []models.RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
			},
		},
		{
			name: "equal",
			args: args{
				left: models.SourceState{
					Name: "left",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				right: models.SourceState{
					Name: "right",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
			},
			want: SourceDiff{
				NameOfLeft:  "left",
				NameOfRight: "right",
				Equal: []models.RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
			},
		},
		{
			name: "diff",
			args: args{
				left: models.SourceState{
					Name: "left",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
				right: models.SourceState{
					Name: "right",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
						{Name: "three", LastCommit: "350"},
						{Name: "four", LastCommit: "450"},
					},
				},
			},
			want: SourceDiff{
				NameOfLeft:  "left",
				NameOfRight: "right",
				Equal: []models.RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
				Diff: []RepoDiff{
					{Name: "three", LastCommitInLeft: "300", LastCommitInRight: "350"},
					{Name: "four", LastCommitInLeft: "400", LastCommitInRight: "450"},
				},
			},
		},
		{
			name: "left with outsiders",
			args: args{
				left: models.SourceState{
					Name: "left",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
				right: models.SourceState{
					Name: "right",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
			},
			want: SourceDiff{
				NameOfLeft:  "left",
				NameOfRight: "right",
				Equal: []models.RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
				MissedInRight: []models.RepoState{
					{Name: "three", LastCommit: "300"},
					{Name: "four", LastCommit: "400"},
				},
			},
		},
		{
			name: "right with outsiders",
			args: args{
				left: models.SourceState{
					Name: "left",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					},
				},
				right: models.SourceState{
					Name: "right",
					Repos: []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					},
				},
			},
			want: SourceDiff{
				NameOfLeft:  "left",
				NameOfRight: "right",
				Equal: []models.RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
				MissedInLeft: []models.RepoState{
					{Name: "three", LastCommit: "300"},
					{Name: "four", LastCommit: "400"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompareRepos(tt.args.left, tt.args.right)

			assert.Equal(t, tt.want, got)
		})
	}
}
