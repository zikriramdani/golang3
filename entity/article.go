package entity

// Article object for REST(CRUD)
type Article struct {
	ID       int    `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required"`
}
