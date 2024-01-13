package interfaces

import "github.com/ashiqsabith123/auth-svc/pkg/domain"

type UserRepo interface {
	FindUser(phone string) (userID uint, err error)
	CreateUser(newUser domain.User) (userID uint, err error)
	SaveUserDetails(userDetails domain.UserDetails) error
	GetUserByID(id uint) (userDetails domain.UserDetails, err error)
	GetUsersByGender(gender string) (userDetails []domain.UserDetails, err error)
}
