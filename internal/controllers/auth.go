package controllers

import (
	"BlacAi/internal/db"
	"BlacAi/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignupAuth(c *gin.Context){

	
	SignupUserDataRaw,exist:= c.Get("SignupBody")

	if !exist{

		c.JSON(http.StatusBadRequest,gin.H{"error":"empty input data"})
		return
	}

	SignupUserData,_:=SignupUserDataRaw.(models.UserSignupInput)


	
	//check for the email in the database 
	user,err:= gorm.G[models.UserDetails](db.DB).Where("Email = ?", SignupUserData.Email).First(c.Request.Context())

	if err!= nil{

		if errors.Is(err,gorm.ErrRecordNotFound){
			c.JSON(http.StatusOK,gin.H{"status":"user not found"})
			return 
		}
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"status":user})


}