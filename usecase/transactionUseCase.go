package usecase

import (
	"net/http"
	"pt-xyz-multifinance/entity/model"
	"pt-xyz-multifinance/entity/request"
	"pt-xyz-multifinance/entity/response"
	"pt-xyz-multifinance/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type TransactionUsecase struct {
	TransactionRepo repositories.TransactionRepository
	UserRepo        repositories.UserRepository
	PinjamanRepo repositories.PinjamanRepository
	TenorRepo repositories.TenorRepository
	ProductRepo repositories.ProductRepository
	
}

func TransactionUsecaseImpl(TransactionRepo repositories.TransactionRepository,
	UserRepo repositories.UserRepository , PinjamanRepo repositories.PinjamanRepository ,
	TenorRepo repositories.TenorRepository , ProductRepo repositories.ProductRepository) *TransactionUsecase {
	return &TransactionUsecase{
		TransactionRepo: TransactionRepo,
		UserRepo:        UserRepo,
		PinjamanRepo: PinjamanRepo,
		TenorRepo: TenorRepo,
		ProductRepo: ProductRepo,
	}
}


func (h *TransactionUsecase) ListTransaction(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userID := userLogin.(jwt.MapClaims)["id"].(float64)

	user , err_user := h.UserRepo.GetUser(int(userID))
	if err_user != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_user.Error(),
		})
	}

	transaction , err := h.TransactionRepo.GetTransactionByIdUser(user.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}


	data := []response.TransactionResponse{}

	for _ ,item := range transaction {
		temp := response.TransactionResponse{
			IdTransaction: item.IdTransaction,
			IdUser: item.IdUser,
			IdProduct: item.IdProduct,
			NomorKontak: item.NomorKontak,
			OTR: item.OTR,
			AdminFee: item.AdminFee,
			JumlahCicilan: item.JumlahCicilan,
			JumlahBunga: item.JumlahBunga,
			NamaAsset: item.NamaAsset,
		}

		data = append(data, temp)
	}
	return c.JSON(http.StatusOK,  map[string]interface{}{
		"data": data,
	})

}


func (h *TransactionUsecase) CreateTransaction(c echo.Context) error {
	request := new(request.TransactionRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	userLogin := c.Get("userLogin")
	userID := userLogin.(jwt.MapClaims)["id"].(float64)

	user , err_user := h.UserRepo.GetUser(int(userID))
	if err_user != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_user.Error(),
		})
	}

	pinjaman , err := h.PinjamanRepo.GetPinjamanByIdUser(user.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	tenor , err := h.TenorRepo.GetTenorByIdTenor(int(pinjaman.IdTenor))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	product , err := h.ProductRepo.GetProductByIdProduct(request.IdProduct)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	admin := (2 * product.Price) / 100
	
	otr := product.Price 

	jumlah_bunga := float64(product.Price) * 0.10 * float64(0.25)
	Jumah_cicilan := (float64(otr) + float64(admin)) / float64(tenor.Tenor)

	if pinjaman.LimitSaldo < otr {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "saldo pinjaman tidak cukup",
		})
	}

	transaction := model.Transaction{
		IdUser: uint(userID),
		IdProduct: request.IdProduct,
		NomorKontak: request.NomorKontak,
		OTR: otr,
		AdminFee: admin,
		JumlahCicilan: int(Jumah_cicilan),
		JumlahBunga: int(jumlah_bunga),
		NamaAsset: product.NameProduct,
	}

	err_trans := h.TransactionRepo.CreateTransaction(transaction)
	if err_trans != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_trans.Error(),
		})
	}

	update_pinjaman := model.Pinjaman{
		IdTenor: pinjaman.IdTenor,
		IdUser: uint(userID),
		LimitSaldo: pinjaman.LimitSaldo - (otr  + int(jumlah_bunga) + int(Jumah_cicilan)),
		
	}
	err_pinjaman := h.PinjamanRepo.UpdatePinjaman(update_pinjaman)
	if err_pinjaman != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_pinjaman.Error(),
		})
	}

	data := response.TransactionResponse {
		IdTransaction: 4958300438543,
		IdUser: uint(userID),
		IdProduct: product.IdProduct,
		NomorKontak: request.NomorKontak,
		OTR: otr,
		AdminFee: admin,
		JumlahCicilan: int(Jumah_cicilan),
		JumlahBunga: int(jumlah_bunga),
		NamaAsset: product.NameProduct,
		TotalPembayaran: float64(otr)  + jumlah_bunga + Jumah_cicilan,
	}
	return c.JSON(http.StatusOK,  map[string]interface{}{
		"data": data,
	})

}
