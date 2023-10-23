package routes

import (
	"financial-management/modules/controllers"
	"financial-management/modules/repositories"
	"financial-management/modules/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Api(router *gin.Engine, db *gorm.DB) {
	// Routes for Category
	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService, ctx)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/categories", categoryController.Index)
		v1.GET("/categories/:id", categoryController.GetByID)
		v1.POST("/categories", categoryController.Create)
		v1.PATCH("/categories/:id", categoryController.Update)
		v1.DELETE("/categories/:id", categoryController.Delete)
	}

	// Routes for Transaction
	transactionRepository := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository)
	transactionController := controllers.NewTransactionController(transactionService, ctx)

	{
		v1.GET("/transactions", transactionController.Index)
		v1.GET("/transactions/:id", transactionController.GetByID)
		v1.POST("/transactions", transactionController.Create)
		v1.PATCH("/transactions/:id", transactionController.Update)
		v1.DELETE("/transactions/:id", transactionController.Delete)
	}
}
