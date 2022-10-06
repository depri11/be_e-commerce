package categories

import (
	"context"
	"time"

	"github.com/depri11/be_e-commerce/domains"
	"github.com/depri11/be_e-commerce/helpers"
)

type service struct {
	repository domains.CategoryRepository
	ctxTimeout time.Duration
}

func NewService(repo domains.CategoryRepository, ctxTimeout time.Duration) *service {
	return &service{repo, ctxTimeout}
}

func (s *service) GetAll(ctx context.Context, params map[string]interface{}) (*helpers.Response, error) {
	result, err := s.repository.GetAll(ctx, params)
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, nil
}

func (s *service) GetById(ctx context.Context, id int) (*helpers.Response, error) {
	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, nil
}
