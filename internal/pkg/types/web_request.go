package types

// LoginRequest Login request body
type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// CreateUserRequest CreateUser request body
type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
}

// EditUserRequest EditUser request body
type EditUserRequest struct {
	UserId    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	Password  string `json:"password"`
}

// CreatePostRequest CreatePost request body
type CreatePostRequest struct {
	UserId           int64  `json:"user_id"`
	ContentText      string `json:"content_text"`
	ContentImagePath string `json:"content_image_path"`
	Visible          bool   `json:"visible"`
}

// EditPostRequest EditPost request body
type EditPostRequest struct {
	ContentText      *string `json:"content_text"`
	ContentImagePath *string `json:"content_image_path"`
	Visible          *bool   `json:"visible"`
}

// CreatePostCommentRequest CreatePostComment request body
type CreatePostCommentRequest struct {
	ContentText string `json:"content_text"`
}
