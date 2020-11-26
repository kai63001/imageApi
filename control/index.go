package control

import (
	"wallbackend/steal"
	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.JSON(200, steal.ListSteal("wall","newest_wallpapers.php",1))
}
