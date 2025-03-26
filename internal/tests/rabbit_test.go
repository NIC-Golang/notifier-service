package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRabbitMQ struct {
	mock.Mock
}

func (m *MockRabbitMQ) Publish(queue string, message []byte) error {
	args := m.Called(queue, message)
	return args.Error(0)
}

func (m *MockRabbitMQ) Consume(queue string) (<-chan []byte, error) {
	args := m.Called(queue)
	return args.Get(0).(chan []byte), args.Error(1)
}

func TestMockRabbitMQ(t *testing.T) {
	mockRabbit := new(MockRabbitMQ)

	queue := "test_queue"
	message := []byte("hello")

	mockRabbit.On("Publish", queue, message).Return(nil)

	err := mockRabbit.Publish(queue, message)
	assert.NoError(t, err)

	mockRabbit.AssertCalled(t, "Publish", queue, message)
}
