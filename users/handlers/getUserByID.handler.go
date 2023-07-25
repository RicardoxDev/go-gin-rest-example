package handlers

import (
	"goRest/context"
	"goRest/users/entity"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	//idRaw := c.Params.ByName("id")
	// id, err := strconv.Atoi(idRaw)
	// if err != nil {
	// 	panic("estalló el parsing del id")
	// }

	user := &entity.User{}
	result := context.DB.First(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"user": user, "status": "no value"})
	}

	c.JSON(200, user)
}
