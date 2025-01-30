package db

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/SthiraPs/DeliveryApp/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=sithira password=@Proxima_ps12@ dbname=postgres sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Check if the "deliveryapp" database exists
    var count int64
    database.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", "deliveryapp").Scan(&count)
    if count == 0 {
        // Create the "deliveryapp" database
        database.Exec("CREATE DATABASE deliveryapp")
        fmt.Println("Database 'deliveryapp' created successfully!")
    }

	// Connect to the "deliveryapp" database
    dsn = "host=localhost user=sithira password=@Proxima_ps12@ dbname=deliveryapp sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

	// Auto Migrate Tables
	db.AutoMigrate(&models.User{}) // Add your models here

	DB = db
	fmt.Println("Database connected successfully!")
}
