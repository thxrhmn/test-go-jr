package models

import "time"

type Todo struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(100);"`
	Description string    `json:"description" gorm:"type:varchar(1000);"`
	File        string    `json:"file"`
	SubTodos    []SubTodo `json:"subTodos" gorm:"foreignKey:TodoID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}