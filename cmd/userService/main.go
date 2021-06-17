package main

import (
	"github.com/labstack/echo/v4"
	"userService/internal/pkg/user/delivery"
)

func setConfig(e *echo.Echo) {
	userDelivery := delivery.UserDelivery{}
	userDelivery.SetRoutersForUser(e)
}

func main() {
	e := echo.New()
	setConfig(e)
	e.Logger.Fatal(e.Start(":8080"))
}
