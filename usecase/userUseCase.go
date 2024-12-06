package usecase

import (
	"fmt"
	"net/http"
	"pt-xyz-multifinance/entity/model"
	"pt-xyz-multifinance/entity/response"
	"pt-xyz-multifinance/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserUsecase struct {
	UserRepo repositories.UserRepository
}

func UserUsecaseImpl(UserRepo repositories.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo}
}

func (h *UserUsecase) UploadProfile(c echo.Context) error {
	fmt.Println("jalan")
	dataFile := c.Get("dataFile").(string)

	userLogin := c.Get("userLogin")
	userID := userLogin.(jwt.MapClaims)["id"].(float64)

	user , err_user := h.UserRepo.GetUser(int(userID))
	if err_user != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_user.Error(),
		})
	}

	user_update := model.Users{
		Id: uint(userID),
		FullName: user.FullName,
		LegalName: user.LegalName,
		NIK: user.NIK,
		Data: user.Data,
		PlaceOfBirth: user.PlaceOfBirth,
		DateOfBirth: user.DateOfBirth,
		Salary: user.Salary,
		IdCardPhoto: user.IdCardPhoto,
		SelfiePhoto: dataFile,
		Email: user.Email,
		Password: user.Password,
	}

	err := h.UserRepo.UpdateUser(user_update)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK,  map[string]interface{}{
		"url": dataFile,
	})

}

func (h UserUsecase) GetProfile(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userID := userLogin.(jwt.MapClaims)["id"].(float64)

	user , err_user := h.UserRepo.GetUser(int(userID))
	if err_user != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_user.Error(),
		})
	}

	data := response.ProfileResponse{
		FullName: user.FullName,
		LegalName: user.LegalName,
		NIK: user.NIK,
		Data: user.Data,
		PlaceOfBirth: user.PlaceOfBirth,
		DateOfBirth: user.DateOfBirth,
		Salary: user.Salary,
		IdCardPhoto: user.IdCardPhoto,
		SelfiePhoto: user.SelfiePhoto,
		Email: user.Email,
	}

	return c.JSON(http.StatusOK,  map[string]interface{}{
		"data": data,
	})
}