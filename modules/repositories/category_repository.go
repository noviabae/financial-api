package repositories

import (
	"financial-management/modules/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() []models.Category
	FindByID(id uint) (*models.Category, error)
	Save(category *models.Category) (*models.Category, error)
	Update(category *models.Category) (*models.Category, error)
	Delete(id uint) error
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

// Delete implements CategoryRepository.
func (cr *CategoryRepositoryImpl) Delete(id uint) error {
	result := cr.db.Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindAll implements CategoryRepository.
func (cr *CategoryRepositoryImpl) FindAll() []models.Category {
	var categories []models.Category
	cr.db.Find(&categories)
	return categories
}

// FindByID implements CategoryRepository.
func (cr *CategoryRepositoryImpl) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := cr.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// Save implements CategoryRepository.
func (cr *CategoryRepositoryImpl) Save(category *models.Category) (*models.Category, error) {
	err := cr.db.Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

// Update implements CategoryRepository.
func (cr *CategoryRepositoryImpl) Update(category *models.Category) (*models.Category, error) {
	err := cr.db.Save(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db}
}
