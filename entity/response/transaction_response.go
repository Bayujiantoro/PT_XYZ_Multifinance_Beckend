package response

import "github.com/google/uuid"

type TransactionResponse struct {
	IdTransaction   uuid.UUID `json:"Id_transaction"`
	IdUser          uint      `json:"id_user"`
	IdProduct       uint      `json:"id_product"`
	NomorKontak     string    `json:"nomor_kontak"`
	OTR             int       `json:"otr"`
	AdminFee        int       `json:"admin_fee"`
	JumlahCicilan   int       `json:"jumlah_cicilan"`
	JumlahBunga     int       `json:"jumlah_bunga"`
	NamaAsset       string    `json:"nama_asset"`
	TotalPembayaran int   `json:"total_pembayaran"`
	JumlahTenor     string    `json:"jumlah_tenor"`
}
