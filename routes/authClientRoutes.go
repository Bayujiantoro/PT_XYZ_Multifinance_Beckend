package routes

import (
	"pt-xyz-multifinance/connection"
	"pt-xyz-multifinance/repositories"
	"pt-xyz-multifinance/usecase"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	userRepo := repositories.UserRepositoryImpl(connection.DB)

	uc := usecase.AuthUsecaseImpl(userRepo)

	e.POST("/register" , uc.RegisterClient)
	e.POST("/login" , uc.LoginClient)
}