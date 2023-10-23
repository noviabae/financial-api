package controllers

import (
	"financial-management/modules/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService services.TransactionService
	ctx                *gin.Context
}

func NewTransactionController(transactionService services.TransactionService, ctx *gin.Context) TransactionController {
	return TransactionController{transactionService, ctx}
}

func (tc *TransactionController) Index(ctx *gin.Context) {
	data, err := tc.transactionService.GetAll()

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

func (tc *TransactionController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := tc.transactionService.GetByID(id)

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

func (tc *TransactionController) Create(ctx *gin.Context) {
	data, err := tc.transactionService.Create(ctx)

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

func (tc *TransactionController) Update(ctx *gin.Context) {
	data, err := tc.transactionService.Update(ctx)

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

func (tc *TransactionController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := tc.transactionService.Delete(id)

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
		"data":   "TRANSACTION DELETED SUCCESSFULLY",
	})

}
