# trasier-client-go

Trasier Go Client ingests tracing data into the Trasier system.

# Installation
You can install the trasier-client-go library with the following command
```go get github.com/trasiercom/trascier-client-go```

# Usage
## 1. Set environemnt variables
The trasier client needs to know your Trasier clientId and the Trasier clientSecret. You can easily store them as environment variables in your program.

```
import (
	"os"
	"github.com/trasiercom/trasier-client-go/pkg"
)

os.Setenv(pkg.TrasierClientIdKey, "YOUR_CLIENTID")
os.Setenv(pkg.TrasierClientSecretKey, "YOUR_CLIENTSECRET")
```

## 2. Create the TrasierClient
The TrasierClient can easily be created with the following line:
```` 
trasierClient := pkg.NewTrasierClient()
```` 

## 3. Time to send some spans

### Create spans
A span is the main entity within the trasier system and can be a regular span in terms of the OpenTracing standard or a single event.
The Span struct is located in the API package. Just use the span struct to create some spans. 

```
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
		IncomingData:              "<xml><data>Incoming Message from GO</data></xml>",
		FinishProcessingTimestamp: 1521755816011,
		EndTimestamp:              1521755816021,
		OutgoingEndpoint:          createEndpoint("Outgoing Endpoint"),
		OutgoingContentType:       api.XML,
		OutgoingData:              "<xml><data>Outgoing message from GO</data></xml>",
	}
	return
}

func createEndpoint(name string) (Enpoint api.Endpoint) {
	Enpoint = api.Endpoint{
		Name:      name,
		IpAddress: "127.0.0.1",
		Port:      "8080",
		Hostname:  "Localhost",
	}
	return
}
```


### Send spans
Finally you can use the client to send spans. 

``` 
accountId := "YOUR_ACCOUNTID"
spaceKey := "YOUR_SPACEKEY"
trasierClient.SendSpans(accountId, spaceKey, spans)
```

# Authentication
As a consumer you don't have to worry about authorization. The trasier client uses OAuth under the hood and takes care of authorization tasks.