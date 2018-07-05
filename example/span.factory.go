package example

import (
	"github.com/trasiercom/trasier-client-go/api"
)

func createSpans() (spans []api.Span) {
	spans = []api.Span{
		createSpan("b368152c-e9de-424c-89ad-aa8f8b1424d5"),
	}
	return
}

func createSpan(id string) (span api.Span) {
	span = api.Span{
		Id:                        id,
		ParentId:                  "",
		TraceId:                   "5823e2a3-6a96-4187-854f-a72a60a91b20",
		ConversationId:            "6f581887-c907-4b31-9288-1e3140c63eb9",
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
