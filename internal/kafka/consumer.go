package kafka

import (
    "context"
    "encoding/json"
    "log"

    "github.com/segmentio/kafka-go"
    "gateway-logger/internal/models"
    "gateway-logger/internal/storage/clickhouse"
)

func StartConsumer(broker, topic string, chWriter *clickhouse.Writer) {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{broker},
        Topic:   topic,
        GroupID: "log-consumer-group",
    })

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Printf("Kafka read error: %v", err)
            continue
        }

        var record models.RequestRecord
        if err := json.Unmarshal(m.Value, &record); err == nil {
            chWriter.Write(record)
        } else {
            log.Printf("Unmarshal error: %v", err)
        }
    }
}
