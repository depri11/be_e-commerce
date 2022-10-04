package usecases

import (
	"github.com/depri11/be_e-commerce/domains"
	"github.com/depri11/be_e-commerce/helpers"
)

type usecases struct {
	repositories domains.UserRepository
}

func NewUsecases(repositories domains.UserRepository) *usecases {
	return &usecases{repositories}
}

func (u *usecases) GetAll() (*helpers.Response, error) {
	result, err := u.repositories.GetAll()
	if err != nil {
		return nil, err
	}

	res := helpers.Response{Status: 200, Message: "Success", Data: result}
	return &res, nil
}
