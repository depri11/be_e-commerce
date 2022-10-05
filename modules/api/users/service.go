package users

import (
	"log"

	"github.com/depri11/be_e-commerce/domains"
	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/input"
	"github.com/depri11/be_e-commerce/models"
)

type service struct {
	repository domains.UserRepository
}

func NewService(repository domains.UserRepository) *service {
	return &service{repository}
}

func (u *service) GetAll() (*helpers.Response, error) {
	result, err := u.repository.GetAll()
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, nil
}

func (u *service) GetByEmail(email string) (*helpers.Response, error) {
	result, err := u.repository.GetByEmail(email)
	if err != nil {
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, err
}

func (u *service) Login(payload *input.UserLoginInput) (*helpers.Response, error) {
	user, err := u.repository.GetByEmail(payload.Email)
	if err != nil {
		return &helpers.Response{Status: 400, Message: "Failed", Data: "Your email or password incorrect!"}, err
	}

	if !helpers.ComparePassword(payload.Password, user.Password) {
		return &helpers.Response{Status: 400, Message: "Failed", Data: "Your email or password incorrect!"}, err
	}

	new := helpers.NewToken(uint(user.ID), user.Username, user.Roles)
	token, err := new.GenerateJWT()
	if err != nil {
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}
	result := map[string]interface{}{
		"token":        token,
		"refreshToken": "",
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, nil
}

func (u *service) Register(payload *input.UserRegisterInput) (*helpers.Response, error) {
	var user models.User
	user.Fullname = payload.Fullname
	user.Username = payload.Username
	hash, err := helpers.HashPassword(payload.Password)
	if err != nil {
		log.Println(err)
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}
	user.Password = hash
	user.Email = payload.Email
	user.Roles = 0
	user.Verify = false

	result, err := u.repository.Register(&user)
	if err != nil {
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, nil
}

func (u *service) Update(email string, payload *input.UserEditProfileInput) (*helpers.Response, error) {
	user, err := u.repository.GetByEmail(email)
	if err != nil {
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}
	user.Fullname = payload.Fullname
	user.Username = payload.Username
	user.Email = payload.Email
	user.Phone = payload.Phone
	if payload.Password != "" {
		hash, err := helpers.HashPassword(payload.Password)
		if err != nil {
			log.Println(err)
			return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
		}
		user.Password = hash
	}
	result, err := u.repository.Update(email, user)
	if err != nil {
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, nil
}
