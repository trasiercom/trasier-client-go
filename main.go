package main

import (
	"./client"
)

func main() {


	trasierClient := client.NewTrasierClient("accountId", "spaceKey")
	trasierClient.SendSpans(nil)

}
