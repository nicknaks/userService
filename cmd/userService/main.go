package main

import (
	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log2 "github.com/labstack/gommon/log"
	"log"
	"userService/internal/pkg/user/delivery"
	"userService/internal/pkg/user/repository"
	"userService/internal/pkg/user/usecase"
)

func initDB() *hare.Database {
	ds, err := disk.New("./data", ".json")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := hare.New(ds)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.CreateTable("users")
	return db
}

func setConfig(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Logger.SetLevel(log2.Lvl(0))

	db := initDB()

	userRepository := repository.UserRepository{DB: *db}
	userUsecase := usecase.UserUsecase{Repository: &userRepository}
	userDelivery := delivery.UserDelivery{Usecase: &userUsecase}
	userDelivery.SetRoutersForUser(e)
}

func main() {
	e := echo.New()
	setConfig(e)
	e.Logger.Fatal(e.Start(":8080"))
}
