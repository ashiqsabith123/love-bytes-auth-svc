package repository

import (
	"errors"

	"github.com/ashiqsabith123/auth-svc/pkg/domain"
	interfaces "github.com/ashiqsabith123/auth-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserRepo struct {
	Postgres *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserRepo{Postgres: db}
}

func (U *UserRepo) FindUser(phone string) (userID uint, err error) {

	query := "SELECT id FROM users WHERE phone=$1"

	if err = U.Postgres.Raw(query, phone).Scan(&userID).Error; err != nil {
		return 0, err
	}

	return userID, nil

}

func (U *UserRepo) CreateUser(newUser domain.User) (userID uint, err error) {

	if err := U.Postgres.Create(&newUser).Error; err != nil {
		return 0, err
	}

	return newUser.ID, nil
}

func (U *UserRepo) SaveUserDetails(userDetails domain.UserDetails) error {

	if err := U.Postgres.Create(&userDetails).Error; err != nil {
		return errors.New("Error while inserting user data" + err.Error())
	}

	return nil
}
