package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/metadex-dao/we3-gaming/controller"
	"github.com/metadex-dao/we3-gaming/middleware"
	"github.com/metadex-dao/we3-gaming/service"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()
	router.Use(middleware.DatabaseConnect())

	var (
		userService    service.UserService       = service.NewUserService()
		userController controller.UserController = controller.NewUserController(userService)
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	})

	users := router.Group("/user")
	{
		users.GET("/", func(ctx *gin.Context) {
			users, err := userController.FindAll(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, users)
			}
		})

		users.POST("/", func(ctx *gin.Context) {
			err := userController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "User submitted successfuly"})
			}
		})
	}

	// init app
	router.Run("0.0.0.0:8080")
}
