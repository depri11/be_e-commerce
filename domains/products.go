package domains

import (
	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/models"
)

type ProductRepository interface {
	GetAll(params map[string]interface{}) (products *models.Products, err error)
	GetByID(id int) (product *models.Product, err error)
}

type ProductService interface {
	GetAll(params map[string]interface{}) (*helpers.Response, error)
	GetById(id int) (*helpers.Response, error)
}
