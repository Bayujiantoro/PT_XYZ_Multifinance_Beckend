package response

import "time"

type LoginResponse struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token"`
}

type ProfileResponse struct {
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	NIK          string    `json:"NIK"`
	Data         string    `json:"data"`
	PlaceOfBirth string    `json:"place_of_birth"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Salary       int       `json:"salary"`
	IdCardPhoto  string    `json:"Id_card_photo"`
	SelfiePhoto  string    `json:"selfie_photo"`
	Email        string    `json:"email"`
}
