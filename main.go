package main

import (
	"github.com/trasiercom/trasier-client-go/internal/authentication"
	"github.com/trasiercom/trasier-client-go/example"
)

func main() {

	clientId := "YOUR_CLIENTID"
	clientSecret := "YOUR_CLIENTSECRET"

	authService := authentication.NewAuthService(clientId, clientSecret)

	println("Getting the token")
	println("=====================")
	authToken := authService.GetToken()
	println("Auth token: " + authToken.AccessToken)
	println("=====================")

	awesomeConsumer := example.ExampleConsumer{}
	awesomeConsumer.SendSpans()
}
