package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanlearn.teach/ginessential/common"
	"oceanlearn.teach/ginessential/model"
	"oceanlearn.teach/ginessential/response"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, "Data validation error, category name is required", nil)
	}

	c.DB.Create(&requestCategory)

	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
