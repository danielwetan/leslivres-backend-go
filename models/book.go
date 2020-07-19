package models

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      int `json:"author"`
	Genre       int `json:"genre"`
	Status      string `json:"status"`
	Img         string `json:"img"`
}

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      int `json:"author" binding:"required"`
	Genre       int `json:"genre" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Img         string `json:"img" binding:"required"`
}

type UpdateBookInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      int `json:"author"`
	Genre       int `json:"genre"`
	Status      string `json:"status"`
	Img         string `json:"img"`
}
