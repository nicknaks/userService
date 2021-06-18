package repository

import (
	"github.com/jameycribbs/hare"
	"userService/internal/pkg/models"
)

type UserRepositoryInterface interface {
	CreateUser(user models.User) (int, error)
	GetUserById(id int) (models.User, error)
	DeleteUser(id int) error
}

type UserRepository struct {
	DB hare.Database
}

func (u *UserRepository) DeleteUser(id int) error {
	return u.DB.Delete("users", id)
}

func (u *UserRepository) CreateUser(user models.User) (int, error) {
	return u.DB.Insert("users", &user)
}

func (u *UserRepository) GetUserById(id int) (models.User, error) {
	user := models.User{}
	err := u.DB.Find("users", id, &user)
	return user, err
}
