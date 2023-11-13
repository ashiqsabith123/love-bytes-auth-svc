package repository

import (
	interfaces "github.com/ashiqsabith123/auth-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserRepo struct {
	Postgres *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserRepo{Postgres: db}
}

func (U *UserRepo) FindUser(phone int) (bool, error) {
	var count int

	query := "SELECT COUNT(*) FROM users WHERE phone=$1"

	if err := U.Postgres.Raw(query, phone).Scan(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil

}

func (U *UserRepo) CreateUser() {

}
