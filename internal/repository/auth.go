package repository

import (
	"BlacAi/internal/db"
	"BlacAi/internal/models"
	"context"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserByEmail(email string)(models.UserDetails,error)
}

type UserGormRepo struct{
	db *gorm.DB
}


//---call this function in main ad give the db.DB as a parameter there 
func NewUserRepo(db *gorm.DB) *UserGormRepo{
	return &UserGormRepo{db:db}
}

//--fetches the user details based on the mail
func (r *UserGormRepo)GetUserByEmail(email string,ctx context.Context)(models.UserDetails,error){

	user,err:= gorm.G[models.UserDetails](db.DB).Where("Email = ?", email).First(ctx)

	if err != nil{
		return models.UserDetails{},err
	}

	return user,nil
}