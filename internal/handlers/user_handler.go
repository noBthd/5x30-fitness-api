package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/5x30-fitness-api/internal/models"
	"github.com/noBthd/5x30-fitness-api/internal/services"
	"golang.org/x/crypto/bcrypt"
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

func CreateUser(c *gin.Context) {
	var user models.User
	var err error

	user.Email = c.Query("email")
	user.Password = c.Query("password")

	if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

	user.Hashed_password, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
	}

	err = services.CreateUser(user)
	if err != nil {
		log.Println("CREATION USER ERROR: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
}
