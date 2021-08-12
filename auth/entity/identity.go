package entity
// Identity represents an authenticated user identity.
type Identity interface {
	// GetUserID returns the user ID.
	GetUserID() string

	GetUserName() string

	GetAddress() string
}

type TokenResBody struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}