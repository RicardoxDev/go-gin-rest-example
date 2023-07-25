package users

import (
	"goRest/users/handlers"

	"github.com/gin-gonic/gin"
)

func UsersController(rg *gin.Engine) {
	usersRoute := rg.Group("/users")

	usersRoute.GET("/", handlers.GetAllUsers)
	usersRoute.GET("/:id", handlers.GetUserByID)
	usersRoute.GET("/byName/:name", handlers.GetUserByName)
	usersRoute.GET("/ping", func(c *gin.Context) { c.JSON(200, "pong") })

	usersRoute.POST("/", handlers.RegisterUser)
}
