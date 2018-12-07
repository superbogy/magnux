package main

import (
    "net/http"
    "fmt"

	"github.com/gin-gonic/gin"
    "github.com/superbogy/magnux/server/controller"
    "github.com/superbogy/magnux/server/middleware"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 10 << 20 // 8 MiB
	// router.Static("/", "./public")
    gin.SetMode(gin.ReleaseMode)
    router.Use(middleware.Auth)
	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})
    router.POST("/upload", controller.Upload)
    router.POST("/login", controller.Login)

    router.Run(":6639")
    fmt.Println("server run on 6639")
}
