package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mashnoor/basic-api-golang/controllers"
)

func UserRouters(router *gin.Engine) {
	userController := controllers.UserController{}

	v1 := router.Group("/v1")
	v1.POST("/add", userController.AddUser)
}
