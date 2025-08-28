package internal

import (
	"BlacAi/internal/db"
	"log"
)

func main() {

	err:= db.InitDb()


	if err!= nil{
		log.Println("there is an error")
	}

	log.Println("DB object is this ",db.DB)

}