package model

import "time"

type CreateUserDBParam struct {
	ID        int        `gorm:"id"`
	Email     string     `gorm:"email"`
	Name      string     `gorm:"name"`
	Phone     string     `gorm:"phone"`
	CreatedAt *time.Time `gorm:"created_at"`
}
