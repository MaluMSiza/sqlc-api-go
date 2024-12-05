package main

import (
	"log"
	"github.com/MaluMSiza/sqlc-api-go/database"
	"github.com/MaluMSiza/sqlc-api-go/config"
	"github.com/MaluMSiza/sqlc-api-go/internal/user"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	queries := database.New(config.DBPool)

	userRepo := user.NewUserRepository(queries)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	e := echo.New()
	// e.Use(middleware.CORS())


	e.POST("/users", userHandler.CreateUserHandler)
	e.GET("/users/:id", userHandler.GetUserHandler)
	e.PUT("/users/:id", userHandler.UpdateUserHandler)
	e.DELETE("/users/:id", userHandler.DeleteUserHandler)

	log.Fatal(e.Start(":8080"))
}