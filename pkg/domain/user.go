package domain

type User struct {
	ID       int `gorm:"primaryKey;unique"`
	FullName string
	Phone    int64
	Username string
	Password string
}
