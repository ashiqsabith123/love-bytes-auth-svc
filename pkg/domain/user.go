package domain

type User struct {
	ID       string `gorm:"primaryKey;autoIncrement"`
	Fullname string `gorm:"not null"`
	Phone    int64  `gorm:"unique;not null"`
	Username string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
}
