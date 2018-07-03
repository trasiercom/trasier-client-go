package authentication

type AuthToken struct {
	AccessToken       string `json:"access_token"`
	ExpiresIn         int16  `json:"expires_in"`
	RefreshExpiresIn  int16  `json:"refresh_expires_in"`
	RefreshToken      string `json:"refresh_token"`
	TokenType         string `json:"token_type"`
	NotBeforePolicy   int16  `json:"not-before-policy"`
	SessionState      string `json:"session_state"`
	Scope             string `json:"scope"`
}

