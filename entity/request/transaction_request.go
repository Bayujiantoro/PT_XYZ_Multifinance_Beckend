package request

import "github.com/google/uuid"

type TransactionRequest struct {
	IdProduct   uint   `json:"id_product"`
	NomorKontak string `json:"nomor_kontak"`
}

type Payment struct {
	Payment       float64 `json:"payment"`
	IdTransaction uuid.UUID `json:"id_transaction"`
}