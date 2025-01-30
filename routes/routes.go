package routes

import (
	"github.com/SthiraPs/DeliveryApp/controllers"
	"github.com/SthiraPs/DeliveryApp/graphql"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
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
		api.POST("/user/signin", controllers.SignInUserController)
		api.POST("/user/register", controllers.CreateUserController)
		api.GET("/users", controllers.GetAllUsersController)
		api.GET("/user/:id", controllers.GetUserByIdController)
		api.PUT("/user/:id", controllers.UpdateUserByIdController)
		api.DELETE("/user/:id", controllers.DeleteUserByIdController)

		// Driver routes
		//api.POST("/driver/signin", controllers.)
		api.POST("/driver/register", controllers.CreateDriverController)
		api.GET("/drivers", controllers.GetAllDriversController)
		api.GET("/driver/:id", controllers.GetDriverByIdController)
		api.PUT("/driver/:id", controllers.UpdateDriverByIdController)
		api.DELETE("/driver/:id", controllers.DeleteDriverByIdController)
	}

	// GraphQL handler
	graphQLHandler := handler.New(&handler.Config{
		Schema:   &graphql.Schema,
		Pretty:   true,
		GraphiQL: true, // Enables GraphiQL UI
	})
	r.POST("/graphql", gin.WrapH(graphQLHandler))
	r.GET("/graphql", gin.WrapH(graphQLHandler)) // For GraphiQL interface

	return r
}
