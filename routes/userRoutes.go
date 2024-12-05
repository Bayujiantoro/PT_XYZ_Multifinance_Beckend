package routes

import (
	"pt-xyz-multifinance/connection"
	"pt-xyz-multifinance/pkg/middleware"
	"pt-xyz-multifinance/repositories"
	"pt-xyz-multifinance/usecase"

	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Group) {
	userRepo := repositories.UserRepositoryImpl(connection.DB)

	uc := usecase.UserUsecaseImpl(userRepo)

	e.POST("/upload/profile/picture", middleware.Auth(middleware.UploadFile(uc.UploadProfile)))
}
