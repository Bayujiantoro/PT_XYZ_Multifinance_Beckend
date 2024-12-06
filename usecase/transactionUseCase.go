package usecase

import (
	"math"
	"net/http"
	"pt-xyz-multifinance/entity/model"
	"pt-xyz-multifinance/entity/request"
	"pt-xyz-multifinance/entity/response"
	"pt-xyz-multifinance/repositories"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TransactionUsecase struct {
	TransactionRepo repositories.TransactionRepository
	UserRepo        repositories.UserRepository
	PinjamanRepo repositories.PinjamanRepository
	TenorRepo repositories.TenorRepository
	ProductRepo repositories.ProductRepository
	PembayaranRepo repositories.PembayaranRepository
	
}

func TransactionUsecaseImpl(TransactionRepo repositories.TransactionRepository,
	UserRepo repositories.UserRepository , PinjamanRepo repositories.PinjamanRepository ,
	TenorRepo repositories.TenorRepository , ProductRepo repositories.ProductRepository, PembayaranRepo repositories.PembayaranRepository) *TransactionUsecase {
	return &TransactionUsecase{
		TransactionRepo: TransactionRepo,
		UserRepo:        UserRepo,
		PinjamanRepo: PinjamanRepo,
		TenorRepo: TenorRepo,
		ProductRepo: ProductRepo,
		PembayaranRepo: PembayaranRepo,
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
			OTR: int(item.OTR),
			AdminFee: int(item.AdminFee),
			JumlahCicilan: int(item.JumlahCicilan),
			JumlahBunga: int(item.JumlahBunga),
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
	Jumah_cicilan := math.Ceil((float64(otr) + float64(admin) + jumlah_bunga) / float64(tenor.Tenor))

	if pinjaman.LimitSaldo < math.Ceil(Jumah_cicilan) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "saldo pinjaman tidak cukup",
		})
	}
	idUuid := uuid.New()

	transaction := model.Transaction{
		IdTransaction: idUuid,
		IdUser: uint(userID),
		IdProduct: request.IdProduct,
		NomorKontak: request.NomorKontak,
		OTR: otr,
		AdminFee: admin,
		JumlahCicilan: Jumah_cicilan,
		JumlahBunga: jumlah_bunga,
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
		LimitSaldo: pinjaman.LimitSaldo - Jumah_cicilan,
		Cicilan: Jumah_cicilan,
		
	}
	err_pinjaman := h.PinjamanRepo.UpdatePinjaman(update_pinjaman)
	if err_pinjaman != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_pinjaman.Error(),
		})
	}

	

	data := response.TransactionResponse {
		IdTransaction: idUuid,
		IdUser: uint(userID),
		IdProduct: product.IdProduct,
		NomorKontak: request.NomorKontak,
		OTR: int(otr),
		AdminFee: int(admin),
		JumlahCicilan: int(Jumah_cicilan),
		JumlahBunga: int(jumlah_bunga),
		NamaAsset: product.NameProduct,
		TotalPembayaran: int(Jumah_cicilan),
		JumlahTenor: strconv.Itoa(tenor.Tenor) + " Bulan",
	}
	return c.JSON(http.StatusOK,  map[string]interface{}{
		"data": data,
	})

}


func (h *TransactionUsecase) PembayaranCicilan(c echo.Context) error {
	request := new(request.Payment)

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
	if request.Payment != pinjaman.Cicilan {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "jumlah pembayaran tidak sesuai",
		})
	}

	pembayaran := model.Pembayaran{
		IdTenor: tenor.IdTenor,
		IdUser: uint(userID),
		Date: time.Now(),
		Payment: request.Payment,
		IdTransaction: request.IdTransaction,
	}

	err_create := h.PembayaranRepo.CreatePembayaran(pembayaran)
	if err_create != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_create.Error(),
		})
	}
	
	return c.JSON(http.StatusOK,  map[string]interface{}{
		"data": pembayaran,
	})

}


func (h *TransactionUsecase) ListPembayaran(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userID := userLogin.(jwt.MapClaims)["id"].(float64)

	user , err_user := h.UserRepo.GetUser(int(userID))
	if err_user != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err_user.Error(),
		})
	}

	pembayaran , err := h.PembayaranRepo.GetPembayaranByIdUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data := []response.PembayaranResponse{}

	for _ , item := range pembayaran {
		temp := response.PembayaranResponse{
			IdPembayaran: item.IdPembayaran,
			IdTenor: item.IdTenor,
			Name: user.FullName,
			Date: item.Date,
			Payment: item.Payment,
		}
		data = append(data, temp)
	}
	return c.JSON(http.StatusOK,  map[string]interface{}{
		"data": data,
	})

}


func (h *TransactionUsecase) ListProduct(c echo.Context) error {
	product , err := h.ProductRepo.GetProduct()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK,  map[string]interface{}{
		"data": product,
	})
}