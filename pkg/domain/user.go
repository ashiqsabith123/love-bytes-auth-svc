package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   int    `gorm:"primaryKey;autoIncrement" `
	Fullname string `gorm:"not null" json:"fullname,omitempty" validate:"required,min=3,max=50"`
	Phone    string `gorm:"unique" json:"phone,omitempty" validate:"required"`
	Username string `gorm:"unique; not null" json:"username,omitempty" validate:"required,min=3,max=50"`
	Password string `gorm:"not null" json:"password,omitempty" validate:"required,min=6,max=50"`
}
