package systemutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareEnvironmentVariables(t *testing.T) {
	type args struct {
		environmentVariables map[string]string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				environmentVariables: map[string]string{},
			},
			want: nil,
		},
		{
			name: "nonempty",
			args: args{
				environmentVariables: map[string]string{
					"KEY_ONE": "value #1",
					"KEY_TWO": "value #2",
				},
			},
			want: []string{"KEY_ONE=value #1", "KEY_TWO=value #2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrepareEnvironmentVariables(tt.args.environmentVariables)

			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
