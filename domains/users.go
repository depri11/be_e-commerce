package domains

import (
	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/input"
	"github.com/depri11/be_e-commerce/models"
)

type UserRepository interface {
	GetAll() (*models.Users, error)
	GetByEmail(email string) (user *models.User, err error)
	Register(user *models.User) (*models.User, error)
	Update(email string, user *models.User) (*models.User, error)
}

type UserService interface {
	GetAll() (*helpers.Response, error)
	GetByEmail(email string) (*helpers.Response, error)
	Register(payload *input.UserRegisterInput) (*helpers.Response, error)
	Update(email string, payload *input.UserEditProfileInput) (*helpers.Response, error)
}
