package input

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterInput struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEditProfileInput struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
}
