package model


type Tenor struct {
	IdTenor uint   `gorm:"primaryKey;autoIncrement"`
	Tenor int `gorm:"type:int"`
	Limit int `gorm:"type:int"`
	IdUser uint `gorm:"type:int"`
}

func (Tenor) TableName() string {
	return "tenor"
}

