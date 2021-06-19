package repository

import (
	"github.com/jameycribbs/hare"
	"userService/internal/pkg/models"
)

func QueryEpisodes(db *hare.Database, queryFn func(u models.User) bool, limit int) ([]models.User, error) {
	var results []models.User
	var err error

	ids, err := db.IDs("users")
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		e := models.User{}

		if err = db.Find("users", id, &e); err != nil {
			return nil, err
		}

		if queryFn(e) {
			results = append(results, e)
		}

		if limit != 0 && limit == len(results) {
			break
		}
	}

	return results, err
}

type UserRepositoryInterface interface {
	CreateUser(user models.User) (int, error)
	GetUserById(id int) (models.User, error)
	DeleteUser(id int) error
	ChangeUser(user models.User) error
	GetAllUsers() ([]models.User, error)
}

type UserRepository struct {
	DB hare.Database
}

func (u *UserRepository) GetAllUsers() ([]models.User, error) {
	return QueryEpisodes(&u.DB, func(u models.User) bool {
		return true
	}, 0)
}

func (u *UserRepository) ChangeUser(user models.User) error {
	return u.DB.Update("users", &user)
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
