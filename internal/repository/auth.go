package repository

import (
	"BlacAi/internal/models"
	"context"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserByEmail(email string,ctx context.Context)(models.UserDetails,error)
	CreateUser(SingUpDetails models.UserSignupInput,ctx context.Context)(models.UserDetails,error)
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

	user,err:= gorm.G[models.UserDetails](r.db).Where("Email = ?", email).First(ctx)

	if err != nil{
		return models.UserDetails{},err
	}


	return user,nil
}

func (r *UserGormRepo)CreateUser(SingUpDetails models.UserSignupInput,ctx context.Context)(models.UserDetails,error){
	
	var user models.UserDetails
	err:= r.db.Transaction(func(tx *gorm.DB) error {
		user= models.UserDetails{
			Email: SingUpDetails.Email, 
			PhoneNumber: SingUpDetails.PhoneNumber, 
			FirstName: SingUpDetails.FirstName, 
			LastName: SingUpDetails.LastName,
		}
		e:=gorm.G[models.UserDetails](tx).Create(ctx, &user)
		if e!=nil{
			return e
		}
		authProvider:=models.AuthProviderDetails{
			UserId: user.ID,
			HashedPassword: SingUpDetails.Password,
			ProviderName: "local",
		}
		e=gorm.G[models.AuthProviderDetails](tx).Create(ctx,&authProvider)
		if e !=nil{
			return e
		}
		return nil
	})

	if err!=nil{
		return models.UserDetails{},err
	}

	return user,nil

}