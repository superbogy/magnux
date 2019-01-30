package main

import (
    "fmt"
    "github.com/spf13/viper"
    "github.com/superbogy/magnux/server/config"
    "github.com/superbogy/magnux/server/models"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "github.com/superbogy/magnux/server/controller"
)

func main() {
    dir, _ := filepath.Abs("./server/config")
    config.InitConfig(dir, "dev")
    connStr := viper.GetString("db.conn")
    fmt.Println(connStr)
    db, err := models.InitConnect(connStr)
    fmt.Println(db, err)
    defer db.Close()
	router := gin.Default()
	router.MaxMultipartMemory = 20 << 2 // 8 MiB
    gin.SetMode(gin.ReleaseMode)
    //router.Use(middleware.Auth)
	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})
	router.GET("/photos", controller.PhotoList)
    router.POST("/upload", controller.Upload)
    router.POST("/login", controller.Login)
    fmt.Println("server run on 6639")
    router.Run(":6639")

}
