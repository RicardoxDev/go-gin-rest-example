package handlers

import (
	"goRest/context"
	"goRest/users/entity"

	"github.com/gin-gonic/gin"
)

func GetUserByName(c *gin.Context) {
	user := &entity.User{Name: c.Params.ByName("name")}
	result := context.DB.First(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"user": user, "status": "no value"})
	}

	c.JSON(200, user)
}
