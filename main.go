package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superbogy/magnux/server/controllers"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 10 << 20 // 8 MiB
	// router.Static("/", "./public")
	gin.SetMode(gin.ReleaseMode)
	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})
	router.POST("/upload", controllers.Upload)

	router.Run(":6639")
}
