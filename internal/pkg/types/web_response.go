package types

import "time"

// MessageResponse return message in response.
type MessageResponse struct {
	Message string `json:"message"`
}

// PostDetailResponse return post detail in response.
type PostDetailResponse struct {
	PostID           int64     `json:"post_id"`
	UserID           int64     `json:"user_id"`
	ContentText      string    `json:"content_text"`
	ContentImagePath string    `json:"content_image_path"`
	Visible          bool      `json:"visible"`
	CreatedTime      time.Time `json:"created_time"`
}
