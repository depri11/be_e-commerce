package models

type Users struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Roles    int    `json:"roles"`
}
