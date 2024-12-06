package usecase

import (
	"net/http"
	"pt-xyz-multifinance/entity/model"
	"pt-xyz-multifinance/entity/request"
	"pt-xyz-multifinance/entity/response"
	"pt-xyz-multifinance/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type PinjamanUsecase struct {
	PinjamanRepo repositories.PinjamanRepository
	UserRepo     repositories.UserRepository
	TenorRepo    repositories.TenorRepository
}

func PinjamanUsecaseImpl(PinjamanRepo repositories.PinjamanRepository,
	UserRepo repositories.UserRepository,
	TenorRepo repositories.TenorRepository) *PinjamanUsecase {
	return &PinjamanUsecase{
		PinjamanRepo: PinjamanRepo,
		UserRepo:     UserRepo,
		TenorRepo:    TenorRepo,
	}
}

func (h *PinjamanUsecase) CreatePinjaman(c echo.Context) error {
	request := new(request.PinjamanRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	userLogin := c.Get("userLogin")
	userID := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err_user := h.UserRepo.GetUser(int(userID))

	if err_user != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_user.Error(),
		})
	}

	tenorCreate := model.Tenor{
		Tenor:  request.Tenor,
		Limit:  float64(request.LimitSaldo),
		IdUser: uint(userID),
	}

	tenor , err := h.TenorRepo.CreateTenor(tenorCreate)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	pinjaman := model.Pinjaman{
		IdTenor: tenor.IdTenor,
		IdUser:     user.Id,
		Date:       time.Now(),
		LimitSaldo: float64(request.LimitSaldo),
	}

	err_create := h.PinjamanRepo.CreatePinjaman(pinjaman)
	if err_create != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_create.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": pinjaman,
	})

}

func (h *PinjamanUsecase) GetPinjaman(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userID := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err_user := h.UserRepo.GetUser(int(userID))

	if err_user != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_user.Error(),
		})
	}

	pinjaman, err := h.PinjamanRepo.GetPinjamanByIdUser(user.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data := response.PinjamanResponse{
		IdTenor:    pinjaman.IdTenor,
		Date:       pinjaman.Date,
		LimitSaldo: int(pinjaman.LimitSaldo),
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
