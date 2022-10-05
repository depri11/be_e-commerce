package products

import (
	"fmt"
	"log"

	"github.com/depri11/be_e-commerce/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(params map[string]interface{}) (products *models.Products, err error) {
	query := r.db
	if params["order_by"] != "" && params["sort_by"] != "" {
		query = query.Order(fmt.Sprintf("%s %s", params["order_by"], params["sort_by"]))
	}

	err = query.Preload("Galleries").Find(&products).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil
}

func (r *repository) GetByID(id int) (product *models.Product, err error) {
	err = r.db.Where("id = ?", id).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}
