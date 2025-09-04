package service

import (
	"BlacAi/internal/models"
	"BlacAi/internal/repository"
	"context"

	"golang.org/x/crypto/bcrypt"
)


type userService struct {
	Repo repository.UserRepo
}

//--call this function in the main after intializing the 
func NewUserService(Repo repository.UserRepo)*userService{
	return &userService{Repo:Repo}
}



func (u *userService)CreateUserAcc(user models.UserSignupInput,ctx context.Context)(bool,error){

	UserDetails,err:= u.Repo.GetUserByEmail(user.Email)

	if err !=nil{
		return false,err
	}else if UserDetails.Email==""{
		return false,nil
	}else{
		return true,nil
	}
}



func PasswordHashing(password string)(string,error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed),err
}