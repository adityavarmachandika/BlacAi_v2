package middleware

import (
	"BlacAi/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//middleware function that checks the biding of the userinput for the signup
func SignupMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {

		var input models.UserSignupInput


		//check wether the data is accurate with respect to the validations in the  models.UserSignupInput
		err:=c.ShouldBindJSON(&input)

		//there are two kind of erros 1) validation error 2) complete body error. 
		if err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			c.Abort()
			return 
		}

		c.Set("SignupBody",input)
		c.Next()
	}
}

//-- login middleware to check the binding
func LoginMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		var input models.UserLoginInput

		err:= c.ShouldBindJSON(&input)


		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			c.Abort()
			return 
		}

		c.Set("LoginBody",input)
		c.Next()
	}
}