package models

type Category struct {
	CategoryID   int    `json:"category_id" gorm:"primaryKey;auto_increment:true;index"`
	CategoryName string `json:"category_name"  gorm:"type:varchar(255)"`
}
