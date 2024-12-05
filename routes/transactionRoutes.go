package routes

import (
	"pt-xyz-multifinance/connection"
	"pt-xyz-multifinance/pkg/middleware"
	"pt-xyz-multifinance/repositories"
	"pt-xyz-multifinance/usecase"

	"github.com/labstack/echo/v4"
)

func TransactionRouter(e *echo.Group) {
	transactionRepo := repositories.TransactionRepositoryImpl(connection.DB)
	userRepo := repositories.UserRepositoryImpl(connection.DB)
	PinjamanRepo := repositories.PinjamanRepositoryImpl(connection.DB)
	TenorRepo := repositories.TenorRepositoryImpl(connection.DB)
	ProductRepo := repositories.ProductRepositoryImpl(connection.DB)
	uc := usecase.TransactionUsecaseImpl(transactionRepo, userRepo , PinjamanRepo , TenorRepo , ProductRepo)

	e.GET("/transaction", middleware.Auth(uc.ListTransaction))
	e.POST("/transaction/create" , middleware.Auth(uc.CreateTransaction))
}