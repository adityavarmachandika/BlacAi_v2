package main

import (
	"BlacAi/internal"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)
	router:=gin.Default()

	internal.Initialize(router)


	err:=router.Run(":4444")

	if err != nil{
		fmt.Printf("there is an error")
		return 
	}
}