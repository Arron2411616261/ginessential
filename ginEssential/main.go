package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "the length of the telephone must be 11"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "the length of the password must be more than 6"})
			return
		}
		if len(name) == 0 {
			name = randomString(10)
		}
		log.Println(name, telephone, password)
		ctx.JSON(200, gin.H{
			"msg": "register successfully",
		})
	})
	r.Run()
}

func randomString(n int) string {
	var letters = []byte("wqyeoriodndjahngklJGIGSFBGFWMNBKSBGMBWIOQPSZVNXMZScxzmhoiewpqjfdnavgcz")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
