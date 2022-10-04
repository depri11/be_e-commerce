package products

import (
	"github.com/depri11/be_e-commerce/domains"
	"github.com/depri11/be_e-commerce/helpers"
)

type service struct {
	repository domains.ProductRepository
}

func NewService(repository domains.ProductRepository) *service {
	return &service{repository}
}

func (s *service) GetAll() (*helpers.Response, int, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, 404, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, 400, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: products}, 200, nil
}
