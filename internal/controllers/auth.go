package controllers

import (
	"BlacAi/internal/models"
	"BlacAi/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{
	service *service.UserService
}


func NewControllerService (service service.UserService)Controller{
	return Controller{service: &service}
}



func (cntr *Controller) SignupAuth(c *gin.Context){

	
	SignupUserDataRaw,exist:= c.Get("SignupBody")
	if !exist{
		c.JSON(http.StatusBadRequest,gin.H{"error":"empty input data"})
		return
	}

	SignupUserData,_:=SignupUserDataRaw.(models.UserSignupInput)
	

	CreatedUserDetails,err:= cntr.service.CreateUserAcc(SignupUserData,c)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"status-user not created":err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"details":CreatedUserDetails})




}











	// ///////////////////////////
	// //check for the email in the database 
	// user,err:= gorm.G[models.UserDetails](db.DB).Where("Email = ?", SignupUserData.Email).First(c.Request.Context())

	// if err!= nil{

	// 	if errors.Is(err,gorm.ErrRecordNotFound){

	// 		user:=models.UserDetails{Email: SignupUserData.Email,PhoneNumber: SignupUserData.PhoneNumber, FirstName: SignupUserData.FirstName, LastName:SignupUserData.LastName}

	// 		//result variable on passing in the syntax will return the id of the record inserted.
	// 		result:=gorm.WithResult()

	// 		err=gorm.G[models.UserDetails](db.DB,result).Create(c.Request.Context(),&user)

	// 		if err !=nil{
	// 			c.JSON(http.StatusInternalServerError,gin.H{"internal error":err.Error()})
	// 			return
	// 		}


	// 		//populating the authprovider table.
	// 		hashedpassword,err:=service.PasswordHashing(SignupUserData.Password)
	// 		if err!=nil{
	// 			c.JSON(http.StatusUnauthorized,gin.H{"status-password not hashed":err.Error()})
	// 		}
	// 		authProvider:=models.AuthProviderDetails{UserId: user.ID,HashedPassword: hashedpassword,ProviderName: "local"}

	// 		err=gorm.G[models.AuthProviderDetails](db.DB,result).Create(c.Request.Context(),&authProvider)

	// 		if err!=nil{
	// 			c.JSON(http.StatusInternalServerError,gin.H{"status-database error":err.Error()})
	// 		}
	// 		c.JSON(http.StatusOK,gin.H{"status-userCreated":user.ID})
	// 		return 
	// 	}
	// 	c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	// 	return
	// }