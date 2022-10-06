package domains

import (
	"context"

	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/models"
)

type ProductRepository interface {
	GetAll(ctx context.Context, params map[string]interface{}) (products *models.Products, err error)
	GetByID(ctx context.Context, id int) (product *models.Product, err error)
}

type ProductService interface {
	GetAll(ctx context.Context, params map[string]interface{}) (*helpers.Response, error)
	GetById(ctx context.Context, id int) (*helpers.Response, error)
}
