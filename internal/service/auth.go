package service

import (
	"BlacAi/internal/models"
	"BlacAi/internal/repository"
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type UserService struct {
	Repo repository.UserRepo
}

//--call this function in the main after intializing the 
func NewUserService(Repo repository.UserRepo)*UserService{
	return &UserService{Repo:Repo}
}



func (u *UserService)CreateUserAcc(user models.UserSignupInput,ctx context.Context)(models.UserDetails,error){

	//--check for the user email. it may already exists
	_,err:= u.Repo.GetUserByEmail(user.Email,ctx)
	
	if err==nil{
		return models.UserDetails{},err
	}

	//-- if no user found we will recive record not found error. so create a new user
	if errors.Is(err,gorm.ErrRecordNotFound){

		hashedPassword,err:=passwordHashing(user.Password); 
		if err!=nil{
			return models.UserDetails{},err
		}

		user.Password=hashedPassword
		CreatedUser,err:= u.Repo.CreateUser(user,ctx)

		if err!=nil{
			return models.UserDetails{},err
		}
		return CreatedUser,nil
	}

	return models.UserDetails{},err

	

}


func (u *UserService)VerifyLogin(user models.UserLoginInput,ctx context.Context)(models.UserDetails,error){

	UserDetails,err:=u.Repo.GetUserByEmail(user.Email,ctx)

	if err!=nil{

		if errors.Is(err,gorm.ErrRecordNotFound){
			return models.UserDetails{},err
		}
		return models.UserDetails{},err
	}

	ProviderDetails,err:=u.Repo.GetProviderById(UserDetails.ID.String(),ctx)

	if err!=nil{
		return models.UserDetails{},err
	}

	IsPasswordValid:=checkPassword(ProviderDetails.HashedPassword,user.Password)

	if IsPasswordValid !=nil{
		return models.UserDetails{},IsPasswordValid
	}
	return UserDetails,nil

}

func (u * UserService)CreateJWT(id string,email string,firstname string)(string,error){

	godotenv.Load("../.env")
	payload:=jwt.MapClaims{
		"id": id,
		"email": email,
		"name": firstname,
		"exp":time.Now().Add(time.Hour * 24).Unix(),
	}

	secret:=os.Getenv("JWT_Secret")

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	tokenStr,err:=token.SignedString([]byte(secret))

	if err!=nil{
		return "",err
	}
	return tokenStr,nil
}

func passwordHashing(password string)(string,error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed),err
}

func checkPassword(hashedPassword string,password string)(error){
	err:=bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	return err
}