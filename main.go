package main

import (
	"./example"
)

func main() {
	awesomeConsumer := example.ExampleConsumer{}
	awesomeConsumer.SendSpans()
}
