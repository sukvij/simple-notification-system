package service

import (
	"context"
	"fmt"
	"log"
	sns_notification "user-notification/sns-notification"
	userModel "user-notification/user/model"
	"user-notification/user/repository"

	"gorm.io/gorm"
)

type UserService struct {
	Db         *gorm.DB
	SnsService *sns_notification.SNSService
	User       *userModel.User
}

func NewUserService(db *gorm.DB, snsService *sns_notification.SNSService, user *userModel.User) *UserService {
	return &UserService{Db: db, SnsService: snsService, User: user}
}

func (service *UserService) CreateUser() (*userModel.User, error) {
	repo := repository.NewUserRepository(service.Db, service.SnsService, service.User)
	res, err1 := repo.CreateUser()
	if err := service.SnsService.PublishUserCreated(context.Background(), service.User); err1 != nil && err != nil {
		log.Printf("Failed to publish SNS message: %v", err)
	} else {
		fmt.Println("successfully published to sns..")
	}
	return res, err1
}

func (service *UserService) GetUserById() (*userModel.User, error) {
	repo := repository.NewUserRepository(service.Db, service.SnsService, service.User)
	return repo.GetUserByID()
}
