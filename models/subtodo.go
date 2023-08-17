package models

import "time"

type SubTodo struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(100);"`
	Description string    `json:"description" gorm:"type:varchar(1000);"`
	File        string    `json:"file"`
	TodoID      int       `json:"todo_id" gorm:"index"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}