package usecase

import (
	"errors"

	"github.com/ashiqsabith123/auth-svc/pkg/domain"
	repo "github.com/ashiqsabith123/auth-svc/pkg/repository/interface"
	usecase "github.com/ashiqsabith123/auth-svc/pkg/usecase/interfaces"
	"github.com/ashiqsabith123/auth-svc/pkg/utils"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) usecase.UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (U *UserUsecase) SignUp(user domain.User) error {

	err := U.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil

}

func (U *UserUsecase) SendOtp(phone string) (string, error) {

	ok, err := U.UserRepo.FindUser(phone)

	if err != nil {
		return "query error", err
	}

	if ok {
		return "", errors.New("user already exist with this phone number")
	}

	resp, err := utils.SendOtp(phone)

	if err != nil {
		return "Failed to send otp", err
	}

	return resp, nil

}
