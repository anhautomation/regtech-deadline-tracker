package contract

type Response struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
	RequestID string      `json:"requestId"`
}
