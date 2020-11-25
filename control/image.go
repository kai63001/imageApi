package control

import "github.com/gin-gonic/gin"

func ImageControl(c *gin.Context){
	id := c.Param("id")
	c.String(200,id)
}
