package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"oceanlearn.teach/ginessential/common"
	"oceanlearn.teach/ginessential/model"
	"oceanlearn.teach/ginessential/util"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
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
		name = util.RandomString(10)
	}
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "this telephone exists"})
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "hase err"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	//返回结果
	log.Println(name, telephone, password)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "register successfully",
	})
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()

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

	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "this user doesn't exist"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "password error"})
		return
	}

	token := "111"

	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "login in successfully",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
