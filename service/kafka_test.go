package service

import (
    "testing"
    "github.com/IBM/sarama"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
)

// MockProducer to simulate Kafka producer
type MockProducer struct {
    mock.Mock
    sarama.SyncProducer
}

func (m *MockProducer) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
    args := m.Called(msg)
    return int32(args.Int(0)), int64(args.Int(1)), args.Error(2)
}

// TestSendMessage tests the SendMessage function of KafkaProducer.
func TestSendMessage(t *testing.T) {
    mockProducer := new(MockProducer)
    mockProducer.On("SendMessage", mock.Anything).Return(0, 0, nil)

    kafkaProducer := &KafkaProducer{Producer: mockProducer}
    err := kafkaProducer.SendMessage("testTopic", "testMessage")

    assert.NoError(t, err)
    mockProducer.AssertExpectations(t)
}

