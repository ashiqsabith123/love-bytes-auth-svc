package interfaces

import "github.com/ashiqsabith123/auth-svc/pkg/domain"

type UserUsecase interface {
	SignUp(user domain.User)
}
