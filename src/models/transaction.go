package models

type Transaction struct {
	ID uint `json:"id" gorm:"primary_key"`
	User int `json:"user"`
	Book int `json:"book"`
	Status int `json:"status"`
}

type CreateTransactioniInput struct {
	User int `json:"user" binding:"required"`
	Book int `json:"book" binding:"required"`
	Status int `json:"status" binding:"required"`
}

type UpdateTransactionInput struct {
	User int `json:"user"`
	Book int `json:"book"`
	Status int `json:"status"`
}