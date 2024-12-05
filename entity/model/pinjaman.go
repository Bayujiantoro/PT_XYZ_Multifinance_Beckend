package model

import "time"

type Pinjaman struct {
	IdTenor    uint      `gorm:"type:int"`
	IdUser     uint      `gorm:"type:int"`
	Date       time.Time `gorm:""`
	LimitSaldo int       `gorm:"type:int"`
}


func (Pinjaman) TableName() string {
	return "pinjaman"
}