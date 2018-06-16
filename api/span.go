package api

type Span struct {
	Id string
	ParentId string
	TraceId string
	ConversationId string
	OperationName string
	Error bool
	StartTimestamp int64
	BeginProcessingTimestamp int64
	IncomingEndpoint Endpoint
	IncomingContentType ContentType
	IncomingData string
	FinishProcessingTimestamp int64
	EndTimestamp int64
	OutgoingEndpoint Endpoint
	OutgoingContentType ContentType
	OutgoingData string
}