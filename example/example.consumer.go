package example

import (
	"fmt"
	"trasier-client-go/pkg"
)

type ExampleConsumer struct {
}

func (exampleConsumer *ExampleConsumer) SendSpans() {
	spans := createSpans()
	trasierClient := pkg.NewTrasierClient()
	errorMessage := trasierClient.SendSpans("776356", "test-11", spans)
	fmt.Print(errorMessage)
}
