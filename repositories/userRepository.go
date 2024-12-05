package repositories

import (
	"pt-xyz-multifinance/entity/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(id_user int) (model.Users, error)
	GetUserByEmail(email string) (model.Users, error)
	CreateUser(module model.Users) error
	UpdateUser(module model.Users) error
}

func UserRepositoryImpl(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser(id_user int) (model.Users, error) {
	user := model.Users{}

	err := r.db.Debug().Model(&model.Users{}).Where(`id = ?`, id_user).Find(&user).Error

	return user, err
}
func (r *repository)GetUserByEmail(email string) (model.Users, error) {
	user := model.Users{}

	err := r.db.Debug().Model(&model.Users{}).Where(`email = ?`, email).Find(&user).Error

	return user, err
}

func (r *repository) CreateUser(module model.Users) error {
	err := r.db.Debug().Exec(`INSERT INTO users ( full_name, legal_name, nik, data, place_of_birth, date_of_birth, salary, id_card_photo, selfie_photo , email , password) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? , ?)
`, module.FullName, module.LegalName, module.NIK, module.Data, module.PlaceOfBirth, module.DateOfBirth, module.Salary, module.IdCardPhoto, module.SelfiePhoto , module.Email , module.Password).Error

	return err
}


func (r *repository) UpdateUser(module model.Users) error {
	err := r.db.Debug().Model(&model.Users{}).Where(`id = ?`, module.Id).UpdateColumns(&module).Error

	return err
}