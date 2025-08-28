package db

import (
	"fmt"
	"os"
	"BlacAi/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb()(*gorm.DB, error){

	godotenv.Load("../.env")

	conn_string:=os.Getenv("DATABASE_URL")
	fmt.Println(conn_string)
	db,err:=gorm.Open(postgres.Open(conn_string))

	if err !=nil {
		return nil,err;
	}
	return db,nil;
}