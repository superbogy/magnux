package controller

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/superbogy/magnux/server/models"
    "strconv"
)

func PhotoList(ctx *gin.Context) {
    page := ctx.Query("page")
    size := ctx.Query("pagesize")
    fmt.Println(page, size)
    pageNumber, _ := strconv.Atoi(page)
    pagesize, _ := strconv.Atoi(size)
    if pagesize < 1 {
        pagesize = 20
    }

    if pageNumber < 1 {
        pageNumber = 1
    }

    list := models.GetPhoto(pageNumber, pagesize)
    ctx.JSON(200, gin.H{
        "message": "ok",
        "code": 0,
        "data": gin.H {
            "page": pageNumber,
            "size": pagesize,
            "list": list,
        },
    })
}
