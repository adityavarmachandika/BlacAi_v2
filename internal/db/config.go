package db

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func InitDb()( error){

	godotenv.Load("../.env")

	conn_string:=os.Getenv("DATABASE_URL")

	var err error
	DB,err=gorm.Open(postgres.Open(conn_string), &gorm.Config{})

	fmt.Println(DB)
	if err !=nil {
		return err;
	}
	return nil
}