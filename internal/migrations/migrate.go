package main

import (
	"BlacAi/internal/db"
	"BlacAi/internal/models"
	"fmt"
	"log"
)


func  main(){
	
	db,err:=db.InitDb()

	if err != nil{
		log.Fatal("db connection failed❌", err)
	}
	err= db.AutoMigrate(&models.UserDetails{}, &models.AuthProviderDetails{})

	if err!= nil{
		log.Fatal("there is an error while migrating", err)
	}

	fmt.Println("migration done'✅")

}