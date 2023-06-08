package model


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`	
}

type LoginPing struct {
	Username string `json:"username"`
}
