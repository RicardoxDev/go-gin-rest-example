package main

import (
	"goRest/users"

	"github.com/gin-gonic/gin"
)

var App *gin.Engine = gin.Default()

func MountRoutes() {
	users.UsersController(App)
}
