package models



type UserSignupInput struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	PhoneNumber string `json:"phonenumber" binding:"required,e164"`
	FirstName string `json:"firstname" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
}

type UserLoginInput struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}