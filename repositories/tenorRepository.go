package repositories

import (
	"pt-xyz-multifinance/entity/model"

	"gorm.io/gorm"
)

type TenorRepository interface {
	CreateTenor(module model.Tenor) error
	GetTenorByIdTenor(id_tenor int) (model.Tenor , error)
}

func TenorRepositoryImpl(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTenor(module model.Tenor) error {
	err := r.db.Debug().Exec(`insert into tenor(tenor , limit) VALUES(? , ?) `).Error

	return err
	
}

func (r *repository) GetTenorByIdTenor(id_tenor int) (model.Tenor , error) {
	data := model.Tenor{}

	err := r.db.Debug().Model(&model.Tenor{}).Where(`id_tenor = ?`, id_tenor).Find(&data).Error

	return data , err
}
