package response

import "time"

type PinjamanResponse struct {
	IdTenor    uint      `json:"id_tenor"`
	Date       time.Time `json:"date"`
	LimitSaldo int       `json:"limit_saldo"`
}

type PembayaranResponse struct {
	IdPembayaran uint      `json:"id_pembayaran"`
	IdTenor      uint      `json:"id_tenor"`
	Name         string      `json:"name"`
	Date         time.Time `json:"date"`
	Payment      float64   `json:"payment"`
}
