package routes

import (
	"BlacAi/internal/controllers"
	"BlacAi/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {

	r.POST("/auth/signup", middleware.SignupMiddleware(),controllers.SignupAuth)
}