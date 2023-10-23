package dto

type CreateCategoryInput struct {
	CategoryName string `json:"category_name"  validate:"required"`
}
