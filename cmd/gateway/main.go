package main

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "os"
    "time"

    "gateway-logger/internal/kafka"
    "gateway-logger/internal/models"
)

type Route struct {
    Path    string `json:"path"`
    Backend string `json:"backend"`
}

func main() {
    routes := loadRoutes()
    writer := kafka.NewWriter("localhost:9092", "gateway-logs")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        for _, route := range routes {
            if r.URL.Path == route.Path {
                start := time.Now()

                bodyBytes, _ := io.ReadAll(r.Body)
                r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

                req, _ := http.NewRequest(r.Method, route.Backend, bytes.NewBuffer(bodyBytes))
                req.Header = r.Header

                resp, err := http.DefaultClient.Do(req)
                if err != nil {
                    http.Error(w, "Gateway error", 500)
                    return
                }

                io.Copy(w, resp.Body)

                logRecord := models.RequestRecord{
                    IPAddress:          r.RemoteAddr,
                    RequestMethod:      r.Method,
                    RequestPath:        r.URL.Path,
                    RequestQuery:       r.URL.RawQuery,
                    RequestHeaders:     "",
                    RequestBody:        string(bodyBytes),
                    ResponseStatusCode: resp.StatusCode,
                    ResponseTime:       time.Since(start).Seconds(),
                    CreatedAt:          time.Now().Format(time.RFC3339),
                }

                msg, _ := json.Marshal(logRecord)
                kafka.WriteMessage(writer, logRecord.IPAddress, string(msg))
                return
            }
        }
        http.NotFound(w, r)
    })

    http.ListenAndServe(":8000", nil)
}

func loadRoutes() []Route {
    data, _ := os.ReadFile("configs/gateway-config.json")
    var routes []Route
    _ = json.Unmarshal(data, &routes)
    return routes
}
