package handlers

import (
	"net/http"

	services "github.com/CBernieJones/scheduling/src/services"
	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {

	userService := services.UserService{}

	users, err := userService.List()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
