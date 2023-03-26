package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"oceanlearn.teach/ginessential/common"
	"oceanlearn.teach/ginessential/dto"
	"oceanlearn.teach/ginessential/model"
	"oceanlearn.teach/ginessential/response"
	"oceanlearn.teach/ginessential/util"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "the length of the telephone must be 11")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "the length of the password must be more than 6")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "this telephone exists")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "hase err")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "system error"})
		log.Printf("token generate error : %v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "register successfully")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()

	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "the length of the telephone must be 11")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "the length of the password must be more than 6")
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

	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "system error"})
		log.Printf("token generate error : %v", err)
		return
	}

	response.Success(ctx, gin.H{"token": token}, "login in successfully")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
