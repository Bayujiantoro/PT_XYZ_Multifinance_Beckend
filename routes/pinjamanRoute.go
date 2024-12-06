package routes

import (
	"pt-xyz-multifinance/connection"
	"pt-xyz-multifinance/pkg/middleware"
	"pt-xyz-multifinance/repositories"
	"pt-xyz-multifinance/usecase"

	"github.com/labstack/echo/v4"
)

func PinjamanRoutes(e *echo.Group) {
	PinjamanRepo := repositories.PinjamanRepositoryImpl(connection.DB)
	UserRepo := repositories.UserRepositoryImpl(connection.DB)
	TenorRepo := repositories.TenorRepositoryImpl(connection.DB)

	uc := usecase.PinjamanUsecaseImpl(PinjamanRepo,UserRepo,TenorRepo )

	e.POST("/pinjaman/create", middleware.Auth(uc.CreatePinjaman))
	e.GET("/pinjaman", middleware.Auth(uc.GetPinjaman))
}