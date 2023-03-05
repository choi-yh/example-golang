package model

import "time"

type User struct {
	ID        int        `gorm:"id"`
	Email     string     `gorm:"email"`
	Name      string     `gorm:"name"`
	Phone     string     `gorm:"phone"`
	CreatedAt *time.Time `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
}

type UserPassword struct {
	ID       int    `gorm:"id"`
	UserID   int    `gorm:"user_id"`
	Password string `gorm:"password"`
}
