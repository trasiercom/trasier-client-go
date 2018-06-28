package example

import (
	"../client"
	"fmt"
)

type ExampleConsumer struct {
}

func (exampleConsumer *ExampleConsumer) SendSpans() {
	spans := createSpans()
	trasierClient := client.NewTrasierClient()
	errorMessage := trasierClient.SendSpans("776356", "test-11", spans)
	fmt.Print(errorMessage)
}
