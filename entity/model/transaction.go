package model

import "github.com/google/uuid"

type Transaction struct {
	IdTransaction uuid.UUID    `gorm:"primaryKey"`
	IdUser        uint   `gorm:"type:int"`
	IdProduct     uint   `gorm:"type:int"`
	NomorKontak   string `gorm:"type:varchar(100)"`
	OTR           float64    `gorm:"type:int"`
	AdminFee      float64    `gorm:"type:int"`
	JumlahCicilan float64    `gorm:"type:int"`
	JumlahBunga   float64    `gorm:"type:int"`
	NamaAsset     string `gorm:"type:varchar(200)"`
}

func (Transaction) TableName() string {
	return "Transaction"
}
