package models

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Roles    int    `json:"roles"`
}

type Users []User
