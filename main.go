package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"goRest/context"
	"goRest/users/entity"
)

func main() {
	context.ConnectToDB()
	context.DB.AutoMigrate(&entity.User{})

	// Routing
	App.GET("/ping", func(c *gin.Context) { c.JSON(200, "pong") })
	MountRoutes()

	// Port and Runnig
	port := 4000
	App.Run(fmt.Sprint(":", port))
}
