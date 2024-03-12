package service

import (
	"github.com/gin-gonic/gin"
	"github.com/metadex-dao/we3-gaming/entity"
	"gorm.io/gorm"
)

func getDB(c *gin.Context) *gorm.DB {
	return c.Keys["DB"].(*gorm.DB)
}

type UserService interface {
	Save(*gin.Context, entity.User) error
	FindAll(*gin.Context) ([]entity.User, error)

	GetAllUserRoles() []string
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Save(ctx *gin.Context, user entity.User) error {
	session := getDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
	err := session.Save(&user).Error
	return err
}

func (service *userService) FindAll(ctx *gin.Context) ([]entity.User, error) {
	var users []entity.User
	err := getDB(ctx).Find(&users).Error
	return users, err
}

func (service *userService) GetAllUserRoles() []string {
	return []string{"guest", "admin"}
}
