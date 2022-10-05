package domains

import (
	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/models"
)

type CategoryRepository interface {
	GetAll(params map[string]interface{}) (category *models.ProductCategories, err error)
	GetById(id int) (category *models.ProductCategorie, err error)
}

type CategoryService interface {
	GetAll(params map[string]interface{}) (*helpers.Response, error)
	GetById(id int) (*helpers.Response, error)
}
