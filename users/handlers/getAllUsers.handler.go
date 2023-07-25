package handlers

import (
	"goRest/context"
	"goRest/users/entity"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users := []*entity.User{}
	result := context.DB.Model(&entity.User{}).Find(&users)

	if result.Error != nil {
		panic("el webo con el get all")
	}

	c.JSON(200, users)
}
