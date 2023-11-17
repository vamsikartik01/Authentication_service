package models

type SignupForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
