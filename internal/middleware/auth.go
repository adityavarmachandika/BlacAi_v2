package middleware

import (
	"BlacAi/internal/models"
	"fmt"
	"net/http"
	"strings"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

func ProtectedRoute()gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }

		parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
            return
        }
		tokenString := parts[1]

		claims,err:=verifyToken(tokenString)

		if err!=nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		 c.Set("user", claims)

        c.Next()
	}
}



func verifyToken(tokenString string)(jwt.MapClaims,error){

	godotenv.Load("../.env")
	secret:=os.Getenv("JWT_Secret")
	var secretKey = []byte(secret)
	token,err:=jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
   })	

   if err !=nil || !token.Valid{
	return nil,fmt.Errorf("token is invalid")
   }

   claims, ok := token.Claims.(jwt.MapClaims);

    if ok{
		return claims,nil
	}

   return nil,fmt.Errorf("no claims from the user")
}