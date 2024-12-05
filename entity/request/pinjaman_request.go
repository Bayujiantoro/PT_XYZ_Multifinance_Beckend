package request

import "time"

type PinjamanRequest struct {
	IdTenor    uint      `json:"id_tenor"`
	Date       time.Time `json:"date"`
	LimitSaldo int       `json:"limit_saldo"`
}
