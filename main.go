package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 10 << 20 // 8 MiB
	router.Static("/", "./public")
	router.GET("/upload", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")
		files := form.File["files"]
	})

	router.Run(":6639")

}
