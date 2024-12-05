package model

type Transaction struct {
	IdTransaction uint   `gorm:"primaryKey;autoIncrement"`
	IdUser uint  `gorm:"type:int"`
	IdProduct uint `gorm:"type:int"`
	NomorKontak string `gorm:"type:varchar(100)"`
	OTR int `gorm:"type:int"`
	AdminFee int `gorm:"type:int"`
	JumlahCicilan int `gorm:"type:int"`
	JumlahBunga int `gorm:"type:int"`
	NamaAsset string `gorm:"type:varchar(200)"`

}

func (Transaction) TableName() string {
	return "Transaction"
}