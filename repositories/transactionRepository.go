package repositories

import (
	"pt-xyz-multifinance/entity/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(module model.Transaction) error
	GetTransactionByIdUser(id_user uint) ([]model.Transaction , error)
}

func TransactionRepositoryImpl(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(module model.Transaction) error {
	err := r.db.Debug().Exec(`insert into Transaction(nomor_kontak, otr , admin_fee , jumlah_cicilan , jumlah_bunga , nama_asset , id_user , id_product) VALUES (? , ? , ? , ?, ?, ? )`, module.NomorKontak , module.OTR , module.AdminFee , module.JumlahCicilan , module.JumlahBunga , module.NamaAsset , module.IdUser , module.IdProduct).Error

	return err
}

func (r *repository ) GetTransactionByIdUser(id_user uint) ([]model.Transaction , error) {
	data := []model.Transaction{}
	err := r.db.Debug().Model(&model.Transaction{}).Where(`id_user = ?`, id_user). Find(&data).Error

	return data , err
}

