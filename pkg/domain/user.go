package domain

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Fullname string `gorm:"not null"`
	Phone    int64  `gorm:"unique"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
