package internal

import (
	"BlacAi/internal/controllers"
	"BlacAi/internal/db"
	"BlacAi/internal/repository"
	"BlacAi/internal/routes"
	"BlacAi/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {

	db,err:= db.InitDb()

	repo:=repository.NewUserRepo(db)

	service :=service.NewUserService(repo)

	controller:=controllers.NewControllerService(*service)
	if err!= nil{
		log.Println("there is an error connceting to db")
		return
	}

	
	routes.Auth(r,controller)
}