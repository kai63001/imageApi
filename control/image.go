package control

import (
	"wallbackend/steal"

	"github.com/gin-gonic/gin"
)

func ImageControl(c *gin.Context){
	id := c.Param("id")
	c.JSON(200, steal.ImageDetailSteal("wall",id))
}
