package categories

import (
	"github.com/depri11/be_e-commerce/domains"
	"github.com/depri11/be_e-commerce/helpers"
)

type service struct {
	repository domains.CategoryRepository
}

func NewService(repo domains.CategoryRepository) *service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetAll(params map[string]interface{}) (*helpers.Response, error) {
	result, err := s.repository.GetAll(params)
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, nil
}

func (s *service) GetById(id int) (*helpers.Response, error) {
	result, err := s.repository.GetById(id)
	if err != nil {
		if err.Error() == "record not found" {
			return &helpers.Response{Status: 404, Message: "Failed", Data: err.Error()}, err
		}
		return &helpers.Response{Status: 400, Message: "Failed", Data: err.Error()}, err
	}

	return &helpers.Response{Status: 200, Message: "Success", Data: result}, nil
}
