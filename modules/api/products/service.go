package products

import (
	"context"
	"time"

	"github.com/depri11/be_e-commerce/domains"
	"github.com/depri11/be_e-commerce/helpers"
)

type service struct {
	repository domains.ProductRepository
	ctxTimeout time.Duration
}

func NewService(repository domains.ProductRepository, ctxTimeout time.Duration) *service {
	return &service{repository, ctxTimeout}
}

func (s *service) GetAll(ctx context.Context, params map[string]interface{}) (*helpers.Response, error) {
	products, err := s.repository.GetAll(ctx, params)
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: products}, nil
}

func (s *service) GetById(ctx context.Context, id int) (*helpers.Response, error) {
	product, err := s.repository.GetByID(ctx, id)
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
