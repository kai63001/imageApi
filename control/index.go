package control

import (
	"wallbackend/steal"
	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.JSON(200, steal.ListSteal("https://wall.alphacoders.com/newest_wallpapers.php?page=1"))
}
