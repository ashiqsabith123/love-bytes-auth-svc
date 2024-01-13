package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement;not null" `
	Phone string `gorm:"unique;not null" `
}

type UserDetails struct {
	gorm.Model
	UserID      uint   `gorm:"unique;not null" `
	User        User   `gorm:"foreignKey:UserID"`
	Fullname    string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Location    string `gorm:"not null"`
	Latitude    string `gorm:"not null"`
	Longitude   string `gorm:"not null"`
	DateOfBirth string `gorm:"not null"`
	Gender      string `gorm:"not null"`
}
