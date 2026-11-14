package response

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	TraceID string `json:"traceId,omitempty"`
	SpanID  string `json:"spanId,omitempty"`
}
