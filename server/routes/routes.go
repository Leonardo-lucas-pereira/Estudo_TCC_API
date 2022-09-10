package routes

import (
	"github.com/Leonardo-lucas-pereira/tcc-api/controllers"
	"github.com/Leonardo-lucas-pereira/tcc-api/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}

		admin := main.Group("user", middlewares.Auth())
		{
			admin.GET("/", controllers.ListUsers)
			admin.GET("/:id", controllers.GetUser)
		}

		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}

	}
	return router
}
