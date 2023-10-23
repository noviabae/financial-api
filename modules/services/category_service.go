package services

import (
	"financial-management/modules/dto"
	"financial-management/modules/models"
	"financial-management/modules/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CategoryService interface {
	GetAll() []models.Category
	GetByID(id uint) (*models.Category, error)
	Create(ctx *gin.Context) (*models.Category, error)
	Update(ctx *gin.Context) (*models.Category, error)
	Delete(id uint) error
}

type CategoryServiceImpl struct {
	categoryRepository repositories.CategoryRepository
}

// Create implements CategoryService.
func (cs *CategoryServiceImpl) Create(ctx *gin.Context) (*models.Category, error) {
	var input dto.CreateCategoryInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	category := models.Category{
		CategoryName: input.CategoryName,
	}

	result, err := cs.categoryRepository.Save(&category)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete implements CategoryService.
func (cs *CategoryServiceImpl) Delete(id uint) error {
	err := cs.categoryRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements CategoryService.
func (cs *CategoryServiceImpl) GetAll() []models.Category {
	return cs.categoryRepository.FindAll()
}

// GetByID implements CategoryService.
func (cs *CategoryServiceImpl) GetByID(id uint) (*models.Category, error) {
	result, err := cs.categoryRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update implements CategoryService.
func (cs *CategoryServiceImpl) Update(ctx *gin.Context) (*models.Category, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UpdateCategoryInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	category := models.Category{
		CategoryID:   id,
		CategoryName: input.CategoryName,
	}

	result, err := cs.categoryRepository.Update(&category)

	if err != nil {
		return nil, err
	}
	return result, nil

}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{categoryRepository}
}
