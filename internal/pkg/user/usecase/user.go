package usecase

import (
	"userService/internal/pkg/models"
	"userService/internal/pkg/user/repository"
)

type UserUsecaseInterface interface {
	CreateUser(user models.User) (models.User, error)
}

type UserUsecase struct {
	Repository repository.UserRepositoryInterface
}

func (u UserUsecase) CreateUser(user models.User) (models.User, error) {
	id, err := u.Repository.CreateUser(user)

	if err != nil {
		return models.User{}, err
	}

	user.Id = id

	return user, nil
}
