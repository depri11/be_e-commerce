package domains

import (
	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/models"
)

type UserRepository interface {
	GetAll() (*models.Users, error)
}

type UserUsecases interface {
	GetAll() (*helpers.Response, error)
}
