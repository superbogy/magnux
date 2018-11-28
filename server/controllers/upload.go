package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}

	files := form.File["files"]
	fmt.Printf("%+v\n", files)
	for index, file := range files {
		fmt.Println(file, index)
	}
}
