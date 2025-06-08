package main

import (
	"log"
	"user-notification/database"
	sns_notification "user-notification/sns-notification"
	userCtrl "user-notification/user/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := database.Init()
	snsSvc := sns_notification.NewSNSService("arn:aws:sns:us-east-1:000000000000:user-notifications", "http://localstack:4566/000000000000/user-queue")
	app := gin.Default()
	(userCtrl.NewUserController(db, snsSvc)).RegisterRoutes(app)
	snsSvc.RegisterSNSRoutes(app)
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
