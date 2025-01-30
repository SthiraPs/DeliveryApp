package routes

import (
	"github.com/SthiraPs/DeliveryApp/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		api.POST("/users", controllers.CreateUser)
		api.GET("/users", controllers.GetUsers)
	}

	return r
}
