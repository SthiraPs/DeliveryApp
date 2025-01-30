package routes

import (
	"github.com/SthiraPs/DeliveryApp/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger documentation route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := r.Group("/api")
	{
		// User routes
		api.POST("/user", controllers.CreateUser)
		api.GET("/users", controllers.GetAllUsers)
		api.GET("/user/:id", controllers.GetUserById)
		api.PUT("/user/:id", controllers.UpdateUserById)
		api.DELETE("/user/:id", controllers.DeleteUserById)

		// Driver routes
		api.POST("/driver", controllers.CreateDriver)
		api.GET("/drivers", controllers.GetAllDrivers)
		api.GET("/driver/:id", controllers.GetDriverById)
		api.PUT("/driver/:id", controllers.UpdateDriverById)
		api.DELETE("/driver/:id", controllers.DeleteDriverById)
	}

	return r
}
