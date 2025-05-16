package main

import (
	"gateway-logger/internal/kafka"
	"gateway-logger/internal/storage/clickhouse"
)

func main() {
	writer := clickhouse.NewWriter("localhost:9000")
	kafka.StartConsumer("localhost:9092", "gateway-logs", writer)
}
