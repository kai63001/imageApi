package main

import (
	"wallbackend/control"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/",control.GetHome)
	r.GET("/image/:id", control.ImageControl)
	r.Run()
}
