package kafka

import (
    "context"
    "log"
    "github.com/segmentio/kafka-go"
)

func NewWriter(broker, topic string) *kafka.Writer {
    return kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{broker},
        Topic:   topic,
        Balancer: &kafka.LeastBytes{},
    })
}

func WriteMessage(writer *kafka.Writer, key, value string) {
    err := writer.WriteMessages(context.Background(), kafka.Message{
        Key:   []byte(key),
        Value: []byte(value),
    })
    if err != nil {
        log.Printf("Kafka write error: %v", err)
    }
}
