package categories

import (
	"context"
	"log"

	"github.com/depri11/be_e-commerce/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context, params map[string]interface{}) (category *models.ProductCategories, err error) {
	query := r.db
	if params["name"] != nil {
		query = query.Where("lower(name) like ?", "%"+params["name"].(string)+"%")
	}

	err = query.Find(&category).Error
	if err != nil {
		log.Println(err)
		return category, err
	}

	return category, nil
}

func (r *repository) GetById(ctx context.Context, id int) (category *models.ProductCategorie, err error) {
	err = r.db.Where("id = ?", id).Find(&category).Error
	if err != nil {
		log.Println(err)
		return category, err
	}

	return category, nil
}
