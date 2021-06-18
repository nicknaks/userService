package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"userService/internal/pkg/models"
	"userService/internal/pkg/user/usecase"
)

type UserDeliveryInterface interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetUserById(c echo.Context) error
	ChangeUserById(c echo.Context) error
	DeleteUser(c echo.Context) error
}

func (u UserDelivery) SetRoutersForUser(e *echo.Echo) {
	e.POST("/users", u.CreateUser)
	e.GET("/users", u.GetAllUsers)
	e.GET("/users/:id", u.GetUserById)
	e.PATCH("/users/:id", u.ChangeUserById)
	e.DELETE("/users/:id", u.DeleteUser)
}

type UserDelivery struct {
	Usecase usecase.UserUsecaseInterface
}

func (u UserDelivery) CreateUser(c echo.Context) error {
	reqUser := new(models.User)
	err := c.Bind(&reqUser)
	if err != nil {
		c.Logger().Info("Create User - " + err.Error())
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	user, err := u.Usecase.CreateUser(*reqUser)
	if err != nil {
		c.Logger().Info("Create User - " + err.Error())
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	return c.JSON(http.StatusOK, user)
}

func (UserDelivery) GetAllUsers(c echo.Context) error {
	panic("implement me")
}

func (UserDelivery) GetUserById(c echo.Context) error {
	panic("implement me")
}

func (UserDelivery) ChangeUserById(c echo.Context) error {
	panic("implement me")
}

func (UserDelivery) DeleteUser(c echo.Context) error {
	panic("implement me")
}
