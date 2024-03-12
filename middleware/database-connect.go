package middleware

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DatabaseConnect() gin.HandlerFunc {
	dsn := url.URL{
		User:     url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", os.Getenv("DB_SERVER"), os.Getenv("DB_PORT")),
		Path:     os.Getenv("DB_NAME"),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		DryRun: false,
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			}),
	})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database: Connection Successful")
	}

	// db.AutoMigrate(&entity.User{})

	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
