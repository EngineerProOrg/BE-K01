package types

type RequestPayload struct {
	SessionID string `json:"sessionId,omitempty" binding:"required"`
}

type LoginPayload struct {
	UserName string `json:"userName,omitempty" binding:"required"`
}
