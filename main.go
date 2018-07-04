package main

import (
	"github.com/trasiercom/trasier-client-go/example"
	"os"
	"github.com/trasiercom/trasier-client-go/pkg"
)

func main() {

	os.Setenv(pkg.TrasierClientIdKey, "YOUR_CLIENTID")
	os.Setenv(pkg.TrasierClientSecretKey, "YOUR_CLIENTSECRET")

	awesomeConsumer := example.ExampleConsumer{}
	awesomeConsumer.SendSpans()
}
