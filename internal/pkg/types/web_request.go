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

// FollowUserRequest FollowUser request body
type FollowUserRequest struct {
	UserId      int64 `json:"user_id"`
	FollowingId int64 `json:"following_id"`
}
