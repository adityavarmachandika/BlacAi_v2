package routes

import (
	"BlacAi/internal/controllers"
	"BlacAi/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine,handler controllers.Controller) {

	r.POST("/auth/signup", middleware.SignupMiddleware(),handler.SignupAuth)

	r.POST("/auth/login",middleware.LoginMiddleware(),handler.LoginAuth)


}












//--sample get function
	// r.GET("/sampleroute", middleware.ProtectedRoute(),func(c *gin.Context) {
	// 	claims,_:=c.Get("user")
	// 	c.JSON(http.StatusOK,gin.H{"status":claims})
	// })