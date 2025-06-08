package repository

import (
	sns_notification "user-notification/sns-notification"
	userModel "user-notification/user/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db         *gorm.DB
	SnsService *sns_notification.SNSService
	User       *userModel.User
}

func NewUserRepository(db *gorm.DB, snsService *sns_notification.SNSService, user *userModel.User) *UserRepository {
	return &UserRepository{Db: db, SnsService: snsService, User: user}
}

func (repository *UserRepository) CreateUser() (*userModel.User, error) {
	if err := repository.Db.Create(repository.User).Error; err != nil {
		return nil, err
	}
	return repository.User, nil
}

func (repository *UserRepository) GetUserByID() (*userModel.User, error) {
	if err := repository.Db.Find(repository.User).Error; err != nil {
		return nil, err
	}
	return repository.User, nil
}
