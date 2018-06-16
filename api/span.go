package api

type Span struct {
	Id                        string      `json:"id"`
	ParentId                  string      `json:"parentId"`
	TraceId                   string      `json:"traceId"`
	ConversationId            string      `json:"conversationId"`
	OperationName             string      `json:"operationName"`
	Error                     bool        `json:"error"`
	StartTimestamp            int64       `json:"startTimestamp"`
	BeginProcessingTimestamp  int64       `json:"beginProcessingTimestamp"`
	IncomingEndpoint          Endpoint    `json:"incomingEndpoint"`
	IncomingContentType       ContentType `json:"incomingContentType"`
	IncomingData              string      `json:"incomingData"`
	FinishProcessingTimestamp int64       `json:"finishProcessingTimestamp"`
	EndTimestamp              int64       `json:"endTimestamp"`
	OutgoingEndpoint          Endpoint    `json:"outgoingEndpoint"`
	OutgoingContentType       ContentType `json:"outgoingContentType"`
	OutgoingData              string      `json:"outgoingData"`
}
