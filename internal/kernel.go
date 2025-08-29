package internal

import (
	"BlacAi/internal/db"
	"BlacAi/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {

	err:= db.InitDb()

	if err!= nil{
		log.Println("there is an error connceting to db")
		return
	}

	routes.Auth(r)



	log.Println("DB object is this ",db.DB)

}