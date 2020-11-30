package control

import (
	"strings"
	"wallbackend/steal"

	"github.com/gin-gonic/gin"
)
func GetHome(c *gin.Context) {
	page := c.Query("page")
	if len(page) == 0 {
		page = "1"
	}
	path := c.Param("path")
	id := c.Query("id")
	if len(id) > 0 {
		path += "?id="+id
	}
	orderby := c.Query("orderby")
	if len(orderby) == 0 {
		orderby = "none"
	}
	typeWall := c.Query("type")
	if len(typeWall) == 0 {
		typeWall = "wall"
	}
	if typeWall == "mobile" {
		path = strings.ReplaceAll(path,"_wallpapers.php","")
	}
	c.JSON(200, steal.ListSteal(typeWall,path,page,orderby))
}
