package dtos

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}