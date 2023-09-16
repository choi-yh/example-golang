package model

import (
	"time"
)

type CreateUserDBParam struct {
	ID        string    `gorm:"id"`
	Email     string    `gorm:"email"`
	Password  string    `gorm:"password"`
	Name      string    `gorm:"name"`
	Phone     string    `gorm:"phone"`
	CreatedAt time.Time `gorm:"created_at"`
}
