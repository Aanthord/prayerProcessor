// /prayerProcessor/service/kafka.go

package service

import (
    "github.com/IBM/sarama"
    "log"
)

// KafkaProducer wraps the Sarama SyncProducer
type KafkaProducer struct {
    Producer sarama.SyncProducer
}

// NewKafkaProducer creates and returns a new KafkaProducer
func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
    config.Producer.Retry.Max = 5                    // Retry up to 5 times to produce the message
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return nil, err
    }

    return &KafkaProducer{Producer: producer}, nil
}

// SendMessage sends a message to the specified topic
func (kp *KafkaProducer) SendMessage(topic, message string) error {
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder(message),
    }

    _, _, err := kp.Producer.SendMessage(msg)
    return err
}

// Close closes the producer connection to Kafka
func (kp *KafkaProducer) Close() error {
    if err := kp.Producer.Close(); err != nil {
        log.Fatalln("Failed to close Kafka producer:", err)
    }
    return nil
}

