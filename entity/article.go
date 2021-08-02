package entity

import (
	"github.com/go-playground/validator/v10"
)

// Article object for REST(CRUD)
type Article struct {
	ID       int    `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required"`
}

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
