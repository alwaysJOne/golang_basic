package main

import (
	"fmt"
	"myBoad/api"
	"myBoad/db"
	"myBoad/router"
)

func main() {

	dbHandler, err := db.ConnectGorm("root:0000@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	apis := api.NewAPI(dbHandler)
	r := router.Router(apis)

	r.Run(":8081")

}
