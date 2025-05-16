package models

type RequestRecord struct {
	IPAddress          string  `json:"ip_address"`
	RequestMethod      string  `json:"method"`
	RequestPath        string  `json:"path"`
	RequestQuery       string  `json:"query"`
	RequestHeaders     string  `json:"headers"`
	RequestBody        string  `json:"body"`
	ResponseStatusCode int     `json:"status_code"`
	ResponseTime       float64 `json:"response_time"`
	CreatedAt          string  `json:"created_at"`
}
