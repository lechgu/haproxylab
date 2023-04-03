package dtos

// AuthenticateRequest ...
type AuthenticateRequest struct {
	APIKey   string `json:"key"`
	Audience string `json:"aud"`
}
