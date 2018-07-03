package main

import (
	"trasier-client-go/example"
	"trasier-client-go/client/authentication"
)

func main() {

	authService := authentication.NewAuthService("trasier-dev_170530_demo1", "MjUxYzRmNDAtMWJiOS00NzhiLWEzMWYtYWQxNWRkZGVmZTg4")

	println("Getting the token")
	println("=====================")
	authToken := authService.GetToken()
	println("Auth token: " + authToken.AccessToken)
	println("=====================")

	awesomeConsumer := example.ExampleConsumer{}
	awesomeConsumer.SendSpans()
}
