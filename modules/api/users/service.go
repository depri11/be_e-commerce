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

func (u *service) GetAll() (*helpers.Response, int, error) {
	result, err := u.repositories.GetAll()
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, 404, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, 400, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, 200, nil
}
