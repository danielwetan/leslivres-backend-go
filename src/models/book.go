package models

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Status      string `json:"status"`
	Img         string `json:"img"`
}

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Img         string `json:"img" binding:"required"`
}

type UpdateBookInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Status      string `json:"status"`
	Img         string `json:"img"`
}
