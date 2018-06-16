package example

import (
	"../api"
)

func createSpans() (spans []api.Span) {
	spans = []api.Span{
		createSpan("8124231b-2902-4ccc-334f-5b6aa06b3c81"),
		createSpan("8124231b-2902-4ccc-334f-5b6aa06b3c82"),
		createSpan("8124231b-2902-4ccc-334f-5b6aa06b3c83"),
	}
	return
}

func createSpan(id string) (span api.Span) {
	span = api.Span{
		Id:                        id,
		ParentId:                  "1",
		TraceId:                   "1",
		ConversationId:            "1",
		OperationName:             "TestOperation",
		Error:                     false,
		StartTimestamp:            1529075380,
		BeginProcessingTimestamp:  1529075380,
		IncomingEndpoint:          createEndpoint("Incoming Endpoint"),
		IncomingContentType:       api.TEXT,
		IncomingData:              "Testdata",
		FinishProcessingTimestamp: 1529075380,
		EndTimestamp:              1529075380,
		OutgoingEndpoint:          createEndpoint("Outgoing Endpoint"),
		OutgoingContentType:       api.TEXT,
		OutgoingData:              "Sample Data",
	}
	return
}
