package clickhouse

import (
    "context"
    "log"
    "github.com/ClickHouse/clickhouse-go/v2"
    "gateway-logger/internal/models"
)

type Writer struct {
    conn clickhouse.Conn
}

func NewWriter(dsn string) *Writer {
    conn, err := clickhouse.Open(&clickhouse.Options{
        Addr: []string{dsn},
        Auth: clickhouse.Auth{
            Database: "default",
            Username: "default",
            Password: "",
        },
    })
    if err != nil {
        panic(err)
    }

    return &Writer{conn: conn}
}

func (w *Writer) Write(record models.RequestRecord) {
    batch, err := w.conn.PrepareBatch(context.Background(), "INSERT INTO logs")
    if err != nil {
        log.Println("Batch error:", err)
        return
    }

    err = batch.Append(
        record.IPAddress,
        record.RequestMethod,
        record.RequestPath,
        record.RequestQuery,
        record.RequestHeaders,
        record.RequestBody,
        record.ResponseStatusCode,
        record.ResponseTime,
        record.CreatedAt,
    )

    if err != nil {
        log.Println("Append error:", err)
        return
    }

    if err = batch.Send(); err != nil {
        log.Println("Send error:", err)
    }
}
