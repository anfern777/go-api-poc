package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/metadex-dao/we3-gaming/entity"
	"github.com/metadex-dao/we3-gaming/service"
)

type UserController interface {
	FindAll(*gin.Context) ([]entity.User, error)
	Save(ctx *gin.Context) error
}

type userController struct {
	service service.UserService
}

var validate *validator.Validate

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (controller *userController) Save(ctx *gin.Context) error {
	var user entity.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	err = validate.Struct(user)
	if err != nil {
		return err
	}

	err = controller.service.Save(ctx, user)
	return err
}

func (controller *userController) FindAll(ctx *gin.Context) ([]entity.User, error) {
	users, err := controller.service.FindAll(ctx)
	return users, err
}
