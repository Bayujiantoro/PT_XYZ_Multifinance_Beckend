package model

type Product struct {
	IdProduct   uint   `gorm:"primaryKey;autoIncrement"`
	NameProduct string `gorm:"type:varchar(100)"`
	Price       float64    `gorm:"type:int"`
}
