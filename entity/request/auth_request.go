package request

import "time"

type RegisterRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`

	FullName     string `json:"full_name" form:"full_name" validate:"required"`
	LegalName    string `json:"legal_name" form:"legal_name" validate:"required"`
	NIK          string `json:"nik" form:"nik" validate:"required"`
	Data         string `json:"data" form:"data" validate:"required"`
	PlaceOfBirth string`json:"place_of_birth" form:"place_of_birth" validate:"required"`
	DateOfBirth  time.Time `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	Salary       int    `json:"salary" form:"salary" validate:"required"`
	IdCardPhoto  string `json:"id_card_photo" form:"id_card_photo" `
	SelfiePhoto  string `json:"selfie_photo" form:"selfie_photo" `
}


type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}