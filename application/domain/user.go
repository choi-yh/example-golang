package domain

import (
	"time"
)

type User struct {
	ID        string     `json:"id" gorm:"id"`
	Email     string     `json:"email" gorm:"email"`
	Password  string     `json:"password" gorm:"password"`
	Name      string     `json:"name" gorm:"name"`
	Phone     string     `json:"phone" gorm:"phone"`
	CreatedAt time.Time  `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

func (u User) TableName() string {
	return "user"
}
