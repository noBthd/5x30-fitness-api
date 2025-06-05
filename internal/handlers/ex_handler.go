package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/5x30-fitness-api/internal/services"
)

func GetExHandler(c *gin.Context) {
    exs, err := services.GetAllExercises()
    if err != nil {
		log.Println("GET_ALL_EX ERROR: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch exercises"})
        return
    }
    c.JSON(http.StatusOK, exs)
}