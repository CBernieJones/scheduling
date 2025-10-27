package api

import (
	"fmt"

	routes "github.com/CBernieJones/scheduling/src/api/routes"
	config "github.com/CBernieJones/scheduling/src/config"
	db "github.com/CBernieJones/scheduling/src/db"
	"github.com/gin-gonic/gin"
)

func Run() error {

	db.Connect(config.Load())

	router := gin.Default()
	routes.RegisterRoutes(router)
	fmt.Println(config.Port())
	return router.Run(config.Port())
}
