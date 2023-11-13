package usecase

import (
	"errors"

	"github.com/ashiqsabith123/auth-svc/pkg/domain"
	repo "github.com/ashiqsabith123/auth-svc/pkg/repository/interface"
	usecase "github.com/ashiqsabith123/auth-svc/pkg/usecase/interfaces"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) usecase.UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (U *UserUsecase) SignUp(user domain.User) error {
	ok, err := U.UserRepo.FindUser(user.Phone)

	if err != nil {
		return err
	}

	if ok {
		return errors.New("User already exist with this phone number")
	}

	err = U.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil

}
