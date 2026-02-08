package model

import (
	"time"
)

type User struct {
	//gorm.Model
	Id        string    `json:"id" gorm:"primaryKey" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
