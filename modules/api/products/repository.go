package products

import (
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

func (r *repository) GetAll() (*models.Products, error) {
	var products models.Products
	err := r.db.Find(&products).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &products, nil
}
