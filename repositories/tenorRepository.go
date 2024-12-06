package repositories

import (
	"pt-xyz-multifinance/entity/model"

	"gorm.io/gorm"
)

type TenorRepository interface {
	CreateTenor(module model.Tenor) (model.Tenor , error)
	GetTenorByIdTenor(id_tenor int) (model.Tenor , error)
	GetTenorByuserId(id_user uint) (model.Tenor , error)
}

func TenorRepositoryImpl(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTenor(module model.Tenor) (model.Tenor , error) {
	err := r.db.Debug().Exec(`insert into tenor(tenor, ` + "`limit`" + ` , id_user) VALUES(?, ?, ?)`, module.Tenor, module.Limit , module.IdUser).Error
    if err != nil {
        return module, err
    }

	tenor , err := r.GetTenorByuserId(module.IdUser)
	if err != nil {
		
		return module , err
	}
	return tenor ,  nil
	
}

func (r *repository) GetTenorByIdTenor(id_tenor int) (model.Tenor , error) {
	data := model.Tenor{}

	err := r.db.Debug().Model(&model.Tenor{}).Where(`id_tenor = ?`, id_tenor).Find(&data).Error

	return data , err
}
func (r *repository) GetTenorByuserId(id_user uint) (model.Tenor , error) {
	data := model.Tenor{}

	err := r.db.Debug().Model(&model.Tenor{}).Where(`id_user = ?`, id_user).Find(&data).Error

	return data , err
}
