package client

const baseURL string = "https://trasier-writer-dev.app.trasier.com/api/accounts"

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