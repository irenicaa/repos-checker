package loader

import (
	"testing"
	"testing/iotest"

	"github.com/irenicaa/repos-checker/models"
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

func TestMultiSource_LoadRepos(t *testing.T) {
	tests := []struct {
		name    string
		sources MultiSource
		want    []models.RepoState
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			sources: MultiSource{},
			want:    []models.RepoState{},
			wantErr: assert.NoError,
		},
		{
			name: "nonempty",
			sources: MultiSource{
				func() Source {
					repos := []models.RepoState{
						{Name: "one", LastCommit: "100"},
						{Name: "two", LastCommit: "200"},
					}

					source := &MockSource{}
					source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

					return source
				}(),
				func() Source {
					repos := []models.RepoState{
						{Name: "three", LastCommit: "300"},
						{Name: "four", LastCommit: "400"},
					}

					source := &MockSource{}
					source.InnerMock.On("LoadRepos").Return(repos, nil).Times(1)

					return source
				}(),
			},
			want: []models.RepoState{
				{Name: "one", LastCommit: "100"},
				{Name: "two", LastCommit: "200"},
				{Name: "three", LastCommit: "300"},
				{Name: "four", LastCommit: "400"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "error",
			sources: MultiSource{
				func() Source {
					source := &MockSource{}
					source.InnerMock.On("Name").Return("source-one").Times(1)
					source.InnerMock.
						On("LoadRepos").
						Return(([]models.RepoState)(nil), iotest.ErrTimeout).
						Times(1)

					return source
				}(),
				&MockSource{},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sources.LoadRepos()

			for _, source := range tt.sources {
				source.(*MockSource).InnerMock.AssertExpectations(t)
			}
			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
