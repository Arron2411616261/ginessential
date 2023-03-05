package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	Telephone string
	Password  string
}

func main() {

	db := InitDB()
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
			name = RandomString(10)
		}
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "exists"})
		}

		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		//返回结果
		log.Println(name, telephone, password)
		ctx.JSON(200, gin.H{
			"msg": "register successfully",
		})
	})
	r.Run()
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	var letters = []byte("wqyeoriodndjahngklJGIGSFBGFWMNBKSBGMBWIOQPSZVNXMZScxzmhoiewpqjfdnavgcz")
	result := make([]byte, n)

	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "123456"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to Db connection, err: " + err.Error())
	}
	_ = db.AutoMigrate(&User{})
	return db
}
