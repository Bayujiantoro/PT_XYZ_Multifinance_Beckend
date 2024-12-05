package repositories

import (
	"pt-xyz-multifinance/entity/model"

	"gorm.io/gorm"
)

type PinjamanRepository interface {
	CreatePinjaman(module model.Pinjaman) error
	UpdatePinjaman(module model.Pinjaman) error 
	GetPinjamanByIdUser(id_user uint) (model.Pinjaman, error)
}

func PinjamanRepositoryImpl(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePinjaman(module model.Pinjaman) error {
	err := r.db.Debug().Exec(`insert into pinjaman(id_tenor ,id_user , date , limit_saldo) VALUES(? , ? , ?)`, module.IdTenor, module.IdUser , module.Date , module.LimitSaldo).Error

	return err
}

func (r *repository) UpdatePinjaman(module model.Pinjaman) error {
	err := r.db.Debug().Exec(`update pinjaman set limit_saldo = ? `, module.LimitSaldo).Where(`id_tenor = ? and id_user = ?`, module.IdTenor , module.IdUser).Error

	return err
}

func (r *repository) GetPinjamanByIdUser(id_user uint) (model.Pinjaman, error) {
	data := model.Pinjaman{}
	err := r.db.Debug().Model(&model.Pinjaman{}).Where(`id_user = ?`, id_user).Find(&data).Error

	return data , err
}