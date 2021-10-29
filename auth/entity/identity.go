package entity

// Identity represents an authenticated user identity.
type Identity interface {
	// GetUserId returns the user ID.
	GetUserId() string

	GetUserName() string

	GetAddress() string

	GetRole() int
}

type TokenResBody struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
