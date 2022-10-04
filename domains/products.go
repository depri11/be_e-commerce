package domains

import (
	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/models"
)

type ProductRepository interface {
	GetAll() (*models.Products, error)
}

type ProductService interface {
	GetAll() (*helpers.Response, int, error)
}
