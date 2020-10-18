package loader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiSource_Name(t *testing.T) {
	tests := []struct {
		name    string
		sources MultiSource
		want    string
	}{
		{
			name:    "empty",
			sources: MultiSource{},
			want:    "",
		},
		{
			name: "non empty",
			sources: []Source{
				func() Source {
					source := &MockSource{}
					source.InnerMock.On("Name").Return("source-one").Times(1)

					return source
				}(),
				func() Source {
					source := &MockSource{}
					source.InnerMock.On("Name").Return("source-two").Times(1)

					return source
				}(),
			},
			want: "source-one|source-two",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sources.Name()

			for _, source := range tt.sources {
				source.(*MockSource).InnerMock.AssertExpectations(t)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
