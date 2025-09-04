package db

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func InitDb()(*gorm.DB, error){
	
	var DB *gorm.DB
	godotenv.Load("../.env")

	conn_string:=os.Getenv("DATABASE_URL")

	var err error
	DB,err=gorm.Open(postgres.Open(conn_string), &gorm.Config{
		PrepareStmt: false,
	})

	fmt.Println("db connected")
	if err !=nil {
		return nil,err;
	}
	return DB,nil
}