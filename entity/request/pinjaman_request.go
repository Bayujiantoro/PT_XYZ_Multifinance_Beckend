package request

import "time"

type PinjamanRequest struct {
	Date       time.Time `json:"date"`
	LimitSaldo int       `json:"limit_saldo"`
	Tenor int `json:"tenor"`
}
