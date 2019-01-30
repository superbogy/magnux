package controller

import (
    "fmt"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gbrlsnchs/jwt"
)

type User struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// Login controller
func Login(ctx *gin.Context) {
    var body User
    if err := ctx.ShouldBindJSON(&body); err != nil {
        ctx.JSON(400, gin.H{
            "code": 10401,
            "message": err.Error(),
        })
        ctx.Abort()
        return
    }

    fmt.Println("%v+\n %v+\n", body.Username, body.Password)
    now := time.Now()
    hs256 := jwt.NewHS256("test123")
    jot := &jwt.JWT{
        Issuer:         "bugbear",
        Subject:        "someone",
        Audience:       "lalal",
        ExpirationTime: now.Add(24 * 30 * 12 * time.Hour).Unix(),
        NotBefore:      now.Add(30 * time.Minute).Unix(),
        IssuedAt:       now.Unix(),
        ID:             "bugbear",
    }

    jot.SetAlgorithm(hs256)
    jot.SetKeyID("123123")
    payload, err := jwt.Marshal(jot)
    if err != nil {
        ctx.JSON(400, gin.H{
            "message": err.Error(),
            "code": "123123",
        })
        ctx.Abort()

        return
    }

    fmt.Println("%v+\n %v+\n", err, payload)
    token, err := hs256.Sign(payload)
    if err != nil {
        // handle error
        ctx.JSON(401, gin.H{
            "message": "invalid token",
            "code": 10401,
        })
        ctx.Abort()
        return
    }
    fmt.Printf("token = %s", token)
    ctx.JSON(200, gin.H{
        "message": "OK",
        "code": "0",
        "data": gin.H{
            "token": token,
        },
    })
}
