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
