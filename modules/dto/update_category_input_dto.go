package dto

type UpdateCategoryInput struct {
	CategoryName string `json:"category_name"  validate:"required"`
}
