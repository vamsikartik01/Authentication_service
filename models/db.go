package models

type Account struct {
	Id                int    `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	PasswordHash      string `json:"password_hash"`
	SaltId            int    `json:"salt_id"`
	CreatedAt         string `json:"date_created"`
	PasswordChangedAt string `json:"password_changed_at"`
}

type Salt struct {
	Id        int    `json:"id"`
	Salt      string `json:"salt"`
	CreatedAt string `json:"created_at"`
}
