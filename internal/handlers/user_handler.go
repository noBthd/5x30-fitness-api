package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/5x30-fitness-api/internal/services"
)

func GetUsersHandler(c *gin.Context) {
    users, err := services.GetAllUsers()
    if err != nil {
		log.Println("GET_ALL_USERS ERROR: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

func UserExists(c *gin.Context) {
	email := c.Query("email")
	
	users, err := services.UserExists(email)
	if err != nil {
		log.Println("USER_EXISTS ERROR: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	c.JSON(http.StatusOK, users)
}


