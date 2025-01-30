package main

import (
	"fmt"
	"log"
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/routes"
	_ "github.com/SthiraPs/DeliveryApp/docs" // Import Swagger docsƒ
)


//@title Delivery App API
//@version 1.0.0
//@description This is a simple delivery app API
//@host localhost:8081
func main() {
	// Connect to DB
	db.ConnectDB()

	// Setup Routes
	router := routes.SetupRouter()

	port := "8081"
	fmt.Println("Server is running on port", port)
	log.Fatal(router.Run(":" + port))
}
