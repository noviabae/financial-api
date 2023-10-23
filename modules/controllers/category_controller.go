package controllers

import (
	"financial-management/modules/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService services.CategoryService
	ctx             *gin.Context
}

func NewCategoryController(categoryService services.CategoryService, ctx *gin.Context) CategoryController {
	return CategoryController{categoryService, ctx}
}

func (cc *CategoryController) Index(ctx *gin.Context) {
	data := cc.categoryService.GetAll()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (cc *CategoryController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := cc.categoryService.GetByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (cc *CategoryController) Create(ctx *gin.Context) {
	data, err := cc.categoryService.Create(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (cc *CategoryController) Update(ctx *gin.Context) {
	data, err := cc.categoryService.Update(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   err,
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (cc *CategoryController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := cc.categoryService.Delete(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   "CATEGORY DELETED SUCCESSFULLY",
	})
}
