package users

import (
	"github.com/depri11/be_e-commerce/domains"
	"github.com/depri11/be_e-commerce/helpers"
)

type service struct {
	repositories domains.UserRepository
}

func NewService(repositories domains.UserRepository) *service {
	return &service{repositories}
}

func (u *service) GetAll() (*helpers.Response, error) {
	result, err := u.repositories.GetAll()
	if err != nil {
		return nil, err
	}

	res := helpers.Response{Status: 200, Message: "Success", Data: result}
	return &res, nil
}
