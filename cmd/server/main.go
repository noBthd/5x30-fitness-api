package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noBthd/5x30-fitness-api/internal/config"
	"github.com/noBthd/5x30-fitness-api/internal/db"
	"github.com/noBthd/5x30-fitness-api/internal/handlers"
)

func main() {
	cfg := config.GetConfig()
	db.ConnectDB(cfg)
	
	router := gin.Default()

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    router.GET("/users", handlers.GetUsersHandler)

	router.GET("/isExists", handlers.UserExists)

    router.Run(":8080")
}
