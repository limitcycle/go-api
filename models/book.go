package models

type Book struct {
	ID          int    `form:"id" json:"id"`
	Author      string `form:"author" json:"author" binding:"required"`
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Status      int    `form:"status" json:"status" binding:"required"`
}
