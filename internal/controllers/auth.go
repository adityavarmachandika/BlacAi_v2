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


func (cntr *Controller)LoginAuth(c *gin.Context){
	LoginUserDataRaw,exists:=c.Get("LoginBody")

	if !exists {
		c.JSON(http.StatusBadRequest,gin.H{"error":"empty input data"})
		return
	}

	LoginUserData:=LoginUserDataRaw.(models.UserLoginInput)

	IsUserVerified:=cntr.service.VerifyLogin(LoginUserData,c)

	if IsUserVerified !=nil{
		c.JSON(http.StatusForbidden,gin.H{"status-no login":IsUserVerified})
		return
	}


	c.JSON(http.StatusOK,gin.H{"status":"user loged in👍"})
}