package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepoStateIndex(t *testing.T) {
	type args struct {
		repos []RepoState
	}

	tests := []struct {
		name string
		args args
		want RepoStateIndex
	}{
		{
			name: "empty",
			args: args{repos: []RepoState{}},
			want: RepoStateIndex{},
		},
		{
			name: "nonempty",
			args: args{
				repos: []RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
			},
			want: RepoStateIndex{
				"one": RepoState{Name: "one", LastCommit: "100"},
				"two": RepoState{Name: "two", LastCommit: "200"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepoStateIndex(tt.args.repos)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFindRepoStateDuplicates(t *testing.T) {
	type args struct {
		reposStates []RepoState
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{reposStates: []RepoState{}},
			want: nil,
		},
		{
			name: "without duplicates",
			args: args{
				reposStates: []RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
			},
			want: nil,
		},
		{
			name: "with one duplicate with same last commits",
			args: args{
				reposStates: []RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "one", LastCommit: "100"},
					{Name: "two", LastCommit: "200"},
				},
			},
			want: []string{"one"},
		},
		{
			name: "with one duplicate with different last commits",
			args: args{
				reposStates: []RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "one", LastCommit: "150"},
					{Name: "two", LastCommit: "200"},
				},
			},
			want: []string{"one"},
		},
		{
			name: "with few duplicates",
			args: args{
				reposStates: []RepoState{
					{Name: "one", LastCommit: "100"},
					{Name: "one", LastCommit: "150"},
					{Name: "two", LastCommit: "200"},
					{Name: "two", LastCommit: "250"},
				},
			},
			want: []string{"one", "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindRepoStateDuplicates(tt.args.reposStates)

			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
