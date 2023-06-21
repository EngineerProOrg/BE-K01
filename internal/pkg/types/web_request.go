package types

// LoginRequest Login request body
type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
