package main

import (
	"BlacAi/internal/db"
	"BlacAi/internal/models"
	"fmt"
	"log"
)


func  main(){
	
	err:=db.InitDb()

	if err != nil{
		log.Fatal("db connection failed❌", err)
	}
	err= db.DB.AutoMigrate(&models.UserDetails{}, &models.AuthProviderDetails{})

	if err!= nil{
		log.Fatal("there is an error while migrating", err)
	}

	fmt.Println("migration done'✅")

}