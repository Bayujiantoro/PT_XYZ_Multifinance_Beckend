package repositories

import (
	"pt-xyz-multifinance/entity/model"

	"gorm.io/gorm"
)

type PembayaranRepository interface {
	CreatePembayaran(module model.Pembayaran) error
	GetPembayaranByIdUser(id_user uint) ([]model.Pembayaran , error)
}

func PembayaranRepositoryImpl(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePembayaran(module model.Pembayaran) error {
	err := r.db.Debug().Exec(`insert into pembayaran(id_tenor, id_user , date , payment , id_transaction  ) VALUES (? , ? , ? , ? , ? )`, module.IdTenor , module.IdUser , module.Date , module.Payment , module.IdTransaction).Error

	return err
}

func (r *repository ) GetPembayaranByIdUser(id_user uint) ([]model.Pembayaran , error) {
	data := []model.Pembayaran{}
	err := r.db.Debug().Model(&model.Pembayaran{}).Where(`id_user = ?`, id_user). Find(&data).Error

	return data , err
}

