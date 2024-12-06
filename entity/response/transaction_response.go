package response


type TransactionResponse struct {
	IdTransaction uint  `json:"Id_transaction"`
	IdUser uint  `json:"id_user"`
	IdProduct uint `json:"id_product"`
	NomorKontak string `json:"nomor_kontak"`
	OTR int `json:"otr"`
	AdminFee int `json:"admin_fee"`
	JumlahCicilan int `json:"jumlah_cicilan"`
	JumlahBunga int `json:"jumlah_bunga"`
	NamaAsset string `json:"nama_asset"`
	TotalPembayaran float64 `json:"total_pembayaran"`
}