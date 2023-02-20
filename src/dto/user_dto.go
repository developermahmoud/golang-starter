package dto

type UserRegisterDTO struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=100"`
	LastName  string `json:"last_name" binding:"required,min=3,max=100"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required,email,unique=users"`
	Locale    string `json:"locale" binding:"required,oneof=en ar"`
}
