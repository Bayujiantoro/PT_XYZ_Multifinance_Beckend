package model

import (
	"time"

	"github.com/google/uuid"
)

type Pembayaran struct {
	IdPembayaran uint      `gorm:"primaryKey;autoIncrement"`
	IdTenor      uint      `gorm:"type:int"`
	IdUser       uint      `gorm:"type:int"`
	Date         time.Time `gorm:""`
	Payment      float64   `gorm:"type:int"`
	IdTransaction uuid.UUID `gorm:""`
}

func (Pembayaran) TableName() string {
	return "pembayaran"
}


