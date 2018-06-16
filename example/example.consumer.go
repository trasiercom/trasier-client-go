package example

import (
	"../client"
	"fmt"
)

type ExampleConsumer struct {
}

func (exampleConsumer *ExampleConsumer) SendSpans() {
	spans := createSpans()
	trasierClient := client.NewTrasierClient("776356", "test-11")
	errorMessage := trasierClient.SendSpans(spans)
	fmt.Print(errorMessage)
}
