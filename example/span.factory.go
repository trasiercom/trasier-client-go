package example

import (
	"github.com/trasiercom/trasier-client-go/api"
)

func createSpans() (spans []api.Span) {
	spans = []api.Span{
		createSpan("8124231b-2902-4ccc-334f-5b6aa06b3c81"),
	}
	return
}

func createSpan(id string) (span api.Span) {
	span = api.Span{
		Id:                        id,
		ParentId:                  "8124231b-2902-4ccc-a480-5b6aa06b3c80",
		TraceId:                   "8124231b-e3b7-4ccc-a480-5b6aa06b3000",
		ConversationId:            "8124231b-e3b7-4ccc-a480-5b6aa06b0000",
		OperationName:             "GET_QUOTE",
		Error:                     false,
		StartTimestamp:            1529075380,
		BeginProcessingTimestamp:  1521755816010,
		IncomingEndpoint:          createEndpoint("Incoming Endpoint"),
		IncomingContentType:       api.XML,
		IncomingData:              "<xml><data>Hello, world IN!</data></xml>",
		FinishProcessingTimestamp: 1521755816011,
		EndTimestamp:              1521755816021,
		OutgoingEndpoint:          createEndpoint("Outgoing Endpoint"),
		OutgoingContentType:       api.XML,
		OutgoingData:              "<xml><data>Hello, world OUT!</data></xml>",
	}
	return
}
