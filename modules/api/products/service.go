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

func (s *service) GetAll(params map[string]interface{}) (*helpers.Response, error) {
	products, err := s.repository.GetAll(params)
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: products}, nil
}

func (s *service) GetById(id int) (*helpers.Response, error) {
	product, err := s.repository.GetByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	if product.ID == 0 {
		return &helpers.Response{Status: 404, Message: "Failed", Data: "Product not found"}, nil
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: product}, nil
}
