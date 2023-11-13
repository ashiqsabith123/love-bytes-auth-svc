package interfaces

import "github.com/ashiqsabith123/auth-svc/pkg/domain"

type UserRepo interface {
	FindUser(phone int64) (bool, error)
	CreateUser(user domain.User) error
}
