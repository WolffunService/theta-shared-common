package entity
// Identity represents an authenticated user identity.
type Identity interface {
	// GetUserID returns the user ID.
	GetUserID() string

	GetUserName() string

	GetAddress() string
}

type TokenResBody struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}