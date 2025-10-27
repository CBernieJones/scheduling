package routes

import (
	"github.com/CBernieJones/scheduling/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", handlers.Users)
	r.GET("/ping", handlers.Pong)
}
