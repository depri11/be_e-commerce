package models

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Fullname  string    `json:"fullname"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Roles     int       `json:"roles"`
	Avatar    string    `json:"avatar"`
	Verify    bool      `json:"verify"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User
