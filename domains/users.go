package domains

import (
	"context"

	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/input"
	"github.com/depri11/be_e-commerce/models"
)

type UserRepository interface {
	GetAll(ctx context.Context) (users *[]models.User, err error)
	GetByEmail(ctx context.Context, email string) (user *models.User, err error)
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, email string, user *models.User) (*models.User, error)
}

type UserService interface {
	GetAll(ctx context.Context) (*helpers.Response, error)
	GetByEmail(ctx context.Context, email string) (*helpers.Response, error)
	Login(ctx context.Context, payload *input.UserLoginInput) (*helpers.Response, error)
	Register(ctx context.Context, payload *input.UserRegisterInput) (*helpers.Response, error)
	Update(ctx context.Context, email string, payload *input.UserEditProfileInput) (*helpers.Response, error)
}
