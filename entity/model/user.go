package model

import "time"

type Users struct {
	Id           uint      `gorm:"primaryKey;autoIncrement"`
	FullName     string    `gorm:"type:varchar(100)"`
	LegalName    string    `gorm:"type:varchar(100)"`
	NIK          string    `gorm:"type:varchar(100)"`
	Data         string    `gorm:"type:varchar(100)"`
	PlaceOfBirth string    `gorm:"type:varchar(100)"`
	DateOfBirth  time.Time `gorm:"type:TIMESTAMP "`
	Salary       int       `gorm:"type:int"`
	IdCardPhoto  string    `gorm:"type:varchar(250)"`
	SelfiePhoto  string    `gorm:"type:varchar(250)"`
	Email        string    `gorm:"type:varchar(250)"`
	Password     string    `gorm:"type:varchar(250)"`
}

func (Users) TableName() string {
	return "Users"
}


   