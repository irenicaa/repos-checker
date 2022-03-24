package loader

import (
	"github.com/irenicaa/repos-checker/v2/models"
	"github.com/stretchr/testify/mock"
)

type MockSource struct {
	InnerMock mock.Mock
}

func (mock *MockSource) Name() string {
	results := mock.InnerMock.Called()
	return results.String(0)
}

func (mock *MockSource) LoadRepos() ([]models.RepoState, error) {
	results := mock.InnerMock.Called()
	return results.Get(0).([]models.RepoState), results.Error(1)
}
