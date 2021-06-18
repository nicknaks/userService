package repository

import (
	"github.com/jameycribbs/hare"
	"userService/internal/pkg/models"
)

type UserRepositoryInterface interface {
	CreateUser(user models.User) (int, error)
}

type UserRepository struct {
	DB hare.Database
}

func (u *UserRepository) CreateUser(user models.User) (int, error) {
	return u.DB.Insert("users", &user)
}
