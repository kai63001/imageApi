package main

import (
	"fmt"
	"os"
	"wallbackend/control"
	"github.com/gin-gonic/gin"
)

func getPort() string {
     var port = os.Getenv("PORT")
     if port == "" {
        port = "8080"
        fmt.Println("No Port In Heroku" + port)
     }
     return ":" + port
}
func main() {
	r := gin.Default()
	r.GET("/p/:path",control.GetHome)
	r.GET("/image/:id", control.ImageControl)
	r.Run(getPort())
}
