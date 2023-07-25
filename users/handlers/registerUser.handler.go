package handlers

import (
	"goRest/context"
	"goRest/users/entity"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	user := &entity.User{}
	if err := c.Bind(user); err != nil {
		panic("explotó el parseo del Register")
	}

	result := context.DB.Create(user)
	if result.Error != nil {
		panic("explotó el Create del Register")
	}

	c.JSON(201, "User Created Succesfully")
}
