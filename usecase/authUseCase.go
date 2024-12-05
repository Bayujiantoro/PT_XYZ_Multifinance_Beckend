package usecase

import (
	"net/http"
	"pt-xyz-multifinance/entity/model"
	"pt-xyz-multifinance/entity/request"
	"pt-xyz-multifinance/entity/response"
	"pt-xyz-multifinance/pkg/bcrypt"
	jwtToken "pt-xyz-multifinance/pkg/jwt"
	"pt-xyz-multifinance/pkg/regexp"
	"pt-xyz-multifinance/repositories"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthUsecase struct {
	UserRepo repositories.UserRepository
}

func AuthUsecaseImpl(UserRepo repositories.UserRepository) *AuthUsecase {
	return &AuthUsecase{UserRepo}
}

func (h *AuthUsecase) RegisterClient(c echo.Context) error {
	request := new(request.RegisterRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	validation := validator.New()

	err := validation.Struct(request)


	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	validEmail := regexp.EmailValidation(request.Email)
	if !validEmail {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "format email invalid",
		})
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	user := model.Users {
		FullName: request.FullName,
		LegalName: request.LegalName,
		NIK: request.NIK,
		Data: request.Data,
		PlaceOfBirth: request.PlaceOfBirth,
		DateOfBirth: request.DateOfBirth,
		Salary: request.Salary,
		IdCardPhoto: "",
		SelfiePhoto: "",
		Email: request.Email,
		Password: password,
	}

	err_user := h.UserRepo.CreateUser(user)
	if err_user != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_user.Error(),
		})
	}

	return c.JSON(http.StatusOK,  map[string]interface{}{
		"data": user,
	})
}

func (h *AuthUsecase) LoginClient(c echo.Context) error {
	request := new(request.LoginRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	validEmail := regexp.EmailValidation(request.Email)
	if !validEmail {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "format email invalid",
		})
	}

	user, err := h.UserRepo.GetUserByEmail(request.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "wrong password",
		})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["name"] = user.FullName
	claims["name"] = user.FullName

	token , errGenerate := jwtToken.GeneratorToken(&claims)
	if errGenerate != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errGenerate.Error(),
		})
	}

	data := response.LoginResponse{
		Email: request.Email,
		Password: user.Password,
		Token: token,
	}

	return c.JSON(http.StatusOK,  map[string]interface{}{
		"data": data,
	})
}