package db

import (
	"fmt"
	"github.com/SthiraPs/DeliveryApp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	if count == 1 {
		database.Exec("DELETE DATABASE deliveryapp")
		fmt.Println("Database 'deliveryapp' deleted successfully!")
	}

	database.Exec("CREATE DATABASE deliveryapp")
	fmt.Println("Database 'deliveryapp' created successfully!")

	// Connect to the "deliveryapp" database
	dsn = "host=localhost user=sithira password=@Proxima_ps12@ dbname=deliveryapp sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate Tables
	AutoMigrate(db)

	DB = db
	fmt.Println("Database connected successfully!")
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Driver{},
		&models.DriverReview{},
		&models.Product{},
		&models.ProductCategory{},
		&models.DeliveryRequest{},
		&models.Branch{},
		&models.Role{},
		&models.Status{},
		&models.VehicleType{},
		&models.Order{},
		&models.OrderItem{},
	)
}
