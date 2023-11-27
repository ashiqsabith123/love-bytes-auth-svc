package usecase

import (
	"errors"

	"github.com/ashiqsabith123/auth-svc/pkg/domain"
	repo "github.com/ashiqsabith123/auth-svc/pkg/repository/interface"
	usecase "github.com/ashiqsabith123/auth-svc/pkg/usecase/interfaces"
	"github.com/ashiqsabith123/auth-svc/pkg/utils"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) usecase.UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (U *UserUsecase) VerifyOtpAndSignUp(req *pb.OtpSignUpReq) (string, int, error) {

	var user domain.User

	copier.Copy(&user, &req)

	err := utils.Validator(user)
	if err != nil {
		return "validation error", 6, err
	}

	status, err := utils.VerifyOtp(req.Phone, req.Otp)
	if err != nil {
		return "Otp verification failed", status, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return "Bycrpt erro", 5, err
	}

	user.Password = string(hash)

	err = U.UserRepo.CreateUser(user)
	if err != nil {
		return "Error in db", status, err
	}

	return "Otp verification sucess", status, nil
}

func (U *UserUsecase) SendOtp(phone string) (string, int, error) {

	ok, err := U.UserRepo.FindUser(phone)

	if err != nil {
		return "query error", 1, err
	}

	if ok {
		return "phone number found", 2, errors.New("user already exist with this phone number")
	}

	resp, err := utils.SendOtp(phone)

	if err != nil {
		return "Failed to send otp", 3, err
	}

	return resp, 4, nil

}
