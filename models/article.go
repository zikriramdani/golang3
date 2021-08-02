package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Article object for REST(CRUD)
type Article struct {
	gorm.Model
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required"`
}

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (b *Article) TableName() string {
	return "article"
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
