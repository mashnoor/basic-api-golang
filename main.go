package main

import (
	"github.com/mashnoor/basic-api-golang/pkg/database"
	"github.com/mashnoor/basic-api-golang/routers"
	"log"
)

func main() {

	if err := database.Connection(); err != nil {
		log.Fatal("database error: %s", err)
	}

	router := routers.Routes()
	router.Run()
	//r := gin.Default()
	//r.GET("/ping", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//r.POST("/add", func(context *gin.Context) {
	//
	//})
	//
	//r.Run()

}
