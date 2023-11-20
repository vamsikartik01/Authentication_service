package models

type SignupForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
