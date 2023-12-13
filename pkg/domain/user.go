package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty" validate:"-"`
	Phone string `gorm:"unique;not null" json:"phone,omitempty"`
}

type UserDetails struct {
	gorm.Model
	UserID      uint   `gorm:"unique;not null" validate:"-"`
	User        User   `gorm:"foreignKey:UserID"`
	Fullname    string `gorm:"not null" json:"fullname,omitempty" validate:"required,max=50"`
	Email       string `gorm:"not null" validate:"required"`
	Location    string `gorm:"not null" validate:"required"`
	DateOfBirth string `gorm:"not null" validate:"required"`
	Gender      string `gorm:"not null" validate:"required"`
}
