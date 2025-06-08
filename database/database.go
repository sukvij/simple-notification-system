package database

import (
	"fmt"
	"log"
	"os"
	"time"
	userModel "user-notification/user/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read environment variables
	dsn := os.Getenv("MYSQL_DSN")
	fmt.Println("dsn is ", dsn)
	var db *gorm.DB
	var err error
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to MySQL (attempt %d/5): %v", i+1, err)
		time.Sleep(5 * time.Second)
	}
	if err1 := db.AutoMigrate(&userModel.User{}); err1 != nil {
		log.Fatalf("Failed to migrate schema: %v", err1)
	}
	return db, err
}
