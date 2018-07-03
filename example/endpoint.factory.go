package example

import (
	"trasier-client-go/api"
)

func createEndpoint(name string) (Enpoint api.Endpoint) {
	Enpoint = api.Endpoint{
		Name:      name,
		IpAddress: "127.0.0.1",
		Port:      "8080",
		Hostname:  "Localhost",
	}
	return
}
