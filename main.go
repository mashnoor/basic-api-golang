package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mashnoor/basic-api-golang/models"
	"github.com/mashnoor/basic-api-golang/pkg/database"
	"log"
)

func main() {

	if err := database.Connection(); err != nil {
		log.Fatal("database error: %s", err)
	}
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/add", func(context *gin.Context) {
		user := new(models.User)

		err := context.ShouldBindJSON(&user)
		if err != nil {
			log.Fatal("Error binding json")
		}

		db := database.GetDB()

		err = db.Create(&user).Error
		if err != nil {
			log.Fatal("Couldn't create user")
		}

		fmt.Println(user.Name)

		context.JSON(200, &user)

	})

	r.Run()

}
