package middleware

import (
    "strings"
    "reflect"
    "bytes"
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/gbrlsnchs/jwt"
)

// Auth middleware
func Auth(ctx *gin.Context) {
    authization := ctx.GetHeader("Authorization")
    if authization == "" {
        ctx.JSON(401, gin.H{
            "message": "authorized required",
            "code": "1401",
        })

        return
    }

    bucket := strings.Split(authization, " ")
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
