package controller

import (
	"net/http"
	"strconv"
	sns_notification "user-notification/sns-notification"
	"user-notification/user/model"
	userService "user-notification/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Db         *gorm.DB
	SnsService *sns_notification.SNSService
}

func NewUserController(db *gorm.DB, snsService *sns_notification.SNSService) *UserController {
	return &UserController{Db: db, SnsService: snsService}
}

func (controller *UserController) RegisterRoutes(r *gin.Engine) {
	r.POST("/users", controller.CreateUser)
	r.GET("/users/:id", controller.GetUserById)
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	user := &model.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service := userService.NewUserService(controller.Db, controller.SnsService, user)
	res, err := service.CreateUser()
	if err != nil {
		ctx.JSON(-1, gin.H{"user creation failed err - ": err})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (controller *UserController) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	service := userService.NewUserService(controller.Db, controller.SnsService, &model.User{ID: uint(id)})
	res, errs := service.GetUserById()
	if errs != nil {
		ctx.JSON(-1, gin.H{"get user by id failed err - ": errs})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
