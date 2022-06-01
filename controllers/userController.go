package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mashnoor/basic-api-golang/models"
	"github.com/mashnoor/basic-api-golang/pkg/registry"
	"log"
)

type UserController struct {
}

func (u *UserController) AddUser(ctx *gin.Context) {
	user := new(models.User)

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal("Error binding json")
	}

	//db := database.GetDB()
	db := registry.GetDB()
	err = db.Create(&user).Error
	if err != nil {
		log.Fatal("Couldn't create user")
	}

	fmt.Println(user.Name)

	ctx.JSON(200, &user)
}
