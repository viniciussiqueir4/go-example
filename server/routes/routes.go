package routes

import (
	"kyc/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		process := main.Group("process")
		{
			process.POST("/", controllers.CreateProduct)
			process.GET("/", controllers.GetProductById)
		}
	}
	return router
}
