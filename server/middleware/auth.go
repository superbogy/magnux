package middleware

import (
    "bytes"
    "fmt"
    "reflect"
    "strings"

    "github.com/gbrlsnchs/jwt"
    "github.com/gin-gonic/gin"
)

// Auth middleware
func Auth(ctx *gin.Context) {
    authorization := ctx.GetHeader("Authorization")
    if authorization == "" {
        ctx.JSON(401, gin.H{
            "message": "authorized required",
            "code": "1401",
        })
        ctx.Abort()
        return
    }

    bucket := strings.Split(authorization, " ")
    if len(bucket) != 2 {
        ctx.JSON(401, gin.H{
            "message": "authorized required",
            "code": "1401",
        })

        return
    }

    fmt.Printf("%v+\n", len(bucket))
    bearer := bucket[0]
    token := bucket[1]
    fmt.Printf("%v+\n %v", bearer, token)

    if bearer != "Bearer" {
        ctx.JSON(401, gin.H{
            "message": "authorized required",
            "code": "1401",
        })

        return
    }

    payload, sig, err := jwt.Parse(token)
    buffer := bytes.NewBuffer(payload)
    fmt.Printf("%v+\n %v+\n %v+\n", sig, buffer.String(), err)
    fmt.Println(reflect.TypeOf(payload), reflect.ValueOf(payload).Kind())
    var jot jwt.JWT
    if err = jwt.Unmarshal(payload, &jot); err != nil {
        ctx.JSON(401, gin.H{
            "message": "authorized required",
            "code": "1401",
        })

        return
    }

    fmt.Println(jot.Algorithm())

    ctx.Next()
}
