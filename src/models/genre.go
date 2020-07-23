package models

type Genre struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type GenreInput struct {
	Name string `json:"name"`
}
