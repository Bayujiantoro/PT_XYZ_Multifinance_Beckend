package repositories

import (
	"pt-xyz-multifinance/entity/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(module model.Product) error 
	GetProductByIdProduct(idProduct uint) (model.Product , error)
}


func ProductRepositoryImpl(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProductByIdProduct(idProduct uint) (model.Product , error) {
	data := model.Product{}

	err := r.db.Debug().Model(&model.Product{}).Where(`id_product = ?`, idProduct).Find(&data).Error

	return data , err
}

func (r *repository) CreateProduct(module model.Product) error {
	err := r.db.Debug().Create(&module).Error

	return err
}