package main

import (
	"BlacAi/internal/db"
	"fmt"

)

func main() {

	db_conn,err:=db.InitDb()

	if err != nil{
		fmt.Println("there is an err", err)
	}
	fmt.Println(db_conn)
}