package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/superbogy/magnux/server/models"
)


type CreateBody struct {
    Name string `json: "name" binding:"required"`
    Title string `json: "title"`
    Description string `json: "description"`
}

func CreateAlbum(ctx *gin.Context) {
    var body CreateBody
    ctx.BindJSON(&body)
    if body.Name == "" {
        ctx.JSON(400, gin.H{
            "message": "miss required params",
            "code": "10400",
        })
        ctx.Abort()
        return
    }

    models.AddAlbum(&body)

    ctx.JSON(200, gin.H{
        "message": "ok",
        "code": 0,
    })
}
