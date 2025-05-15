package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/5x30-fitness-api/internal/services"
)

func GetUsersHandler(c *gin.Context) {
    users, err := services.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}