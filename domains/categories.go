package domains

import (
	"context"

	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/models"
)

type CategoryRepository interface {
	GetAll(ctx context.Context, params map[string]interface{}) (category []*models.ProductCategorie, err error)
	GetById(ctx context.Context, id int) (category *models.ProductCategorie, err error)
}

type CategoryService interface {
	GetAll(ctx context.Context, params map[string]interface{}) (*helpers.Response, error)
	GetById(ctx context.Context, id int) (*helpers.Response, error)
}
