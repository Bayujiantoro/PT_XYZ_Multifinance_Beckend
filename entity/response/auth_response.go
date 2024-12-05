package response


type LoginResponse struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token string `json:"token"`
}