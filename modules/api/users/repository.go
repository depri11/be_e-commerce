package users

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
	return &repository{db}
}

func (r *repository) GetAll(ctx context.Context) (users *[]models.User, err error) {
	err = r.db.Order("id desc").Find(&users).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (user *models.User, err error) {
	err = r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (r *repository) Register(ctx context.Context, user *models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (r *repository) Update(ctx context.Context, email string, user *models.User) (*models.User, error) {
	err := r.db.Save(&user).Where("email = ?", email).Error
	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}
