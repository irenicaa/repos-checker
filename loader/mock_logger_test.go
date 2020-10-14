package loader

import "github.com/stretchr/testify/mock"

type MockLogger struct {
	InnerMock mock.Mock
}

func (mock *MockLogger) Printf(format string, arguments ...interface{}) {
	mock.InnerMock.Called(format, arguments)
}
