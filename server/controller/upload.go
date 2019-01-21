package controller

import (
    "fmt"
    "reflect"

    "github.com/gin-gonic/gin"
    "github.com/upyun/go-sdk/upyun"
)

// Upload upload file to upyun
func Upload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}

    up := upyun.NewUpYun(&upyun.UpYunConfig{
        Bucket:   "knight-test",)
        Operator: "devil",
        Password: "knight123",
    })
	files := form.File["files"]
	fmt.Printf("%+v\n", files)
	for index, file := range files {
        fmt.Println(index, file.Filename)
        fmt.Println("type:", reflect.TypeOf(file))
        src, err := file.Open()
        if err != nil {
            return;
        }

        fmt.Println("type:", reflect.TypeOf(src))
        fmt.Println(up.Put(&upyun.PutObjectConfig{
            Path:      "/demo.log",
            Reader: src,
        }))
    }






}
