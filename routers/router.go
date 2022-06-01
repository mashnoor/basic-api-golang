package routers

import (
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	RegisterRoutes(router)

	return router

}
