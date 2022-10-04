package repositories

import (
	"github.com/depri11/be_e-commerce/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() (*models.Users, error) {
	var users models.Users
	err := r.db.Order("id desc").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}
