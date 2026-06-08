package tokenx

import "fmt"

const (
	TokenTypeBearer = "Bearer"
	TokenTypeSign   = "Sign"
)

var (
	ErrTokenEmpty   = fmt.Errorf("token is empty")
	ErrTokenInvalid = fmt.Errorf("token is invalid")
	ErrTokenExpired = fmt.Errorf("token is expired")
)

type Token struct {
	TokenType         string `json:"token_type"`
	AccessToken       string `json:"access_token"`
	ExpiresIn         int64  `json:"expires_in"`
	RefreshToken      string `json:"refresh_token"`
	RefreshExpiresIn  int64  `json:"refresh_expires_in"`
	RefreshExpiresAt  int64  `json:"refresh_expires_at"`
}
