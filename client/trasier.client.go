package client

const baseURL string = "http://localhost:8081/api/accounts"

type TrasierClient struct {
	AccountId string
	SpaceKey string
}

func NewTrasierClient(accountId string, spaceKey string) *TrasierClient {
	return &TrasierClient{
		AccountId:accountId,
		SpaceKey: spaceKey,
	}
}