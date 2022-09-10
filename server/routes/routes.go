package routes

import (
	"github.com/Leonardo-lucas-pereira/tcc-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
			user.GET("/:id", controllers.GetUser)
			user.GET("/", controllers.ListUsers)
		}

	}
	return router
}
