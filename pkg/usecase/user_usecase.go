package usecase

import (
	"github.com/ashiqsabith123/auth-svc/pkg/domain"
	repo "github.com/ashiqsabith123/auth-svc/pkg/repository/interface"
	usecase "github.com/ashiqsabith123/auth-svc/pkg/usecase/interfaces"
	"github.com/ashiqsabith123/auth-svc/pkg/utils"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
	"github.com/jinzhu/copier"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) usecase.UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (U *UserUsecase) SendOtp(phone string) error {

	err := utils.SendOtp(phone)

	if err != nil {
		return err
	}

	return nil

}

func (U *UserUsecase) VerifyOtpAndAuth(req *pb.VerifyOtpReq) (string, bool, int, error) {

	var userID uint
	var newUser domain.User
	userFound := true

	resp, err := utils.VerifyOtp(req.Phone, req.Otp)
	if err != nil {
		return "", false, resp, err
	}

	userID, err = U.UserRepo.FindUser(req.Phone)
	if err != nil {
		return "", false, 500, err
	}

	if userID == 0 {
		userFound = false
		newUser.Phone = req.Phone
		userID, err = U.UserRepo.CreateUser(newUser)
		if err != nil {
			return "", false, 500, err
		}
	}

	token, err := utils.CreateJWTToken(userID)
	if err != nil {
		return "", false, 500, err
	}

	return token, userFound, 200, nil

}

func (U *UserUsecase) SaveUserDetails(req *pb.UserDetailsReq) error {
	var userDetails domain.UserDetails

	copier.Copy(&userDetails, req)

	err := U.UserRepo.SaveUserDetails(userDetails)

	if err != nil {
		return err
	}

	return nil
}

func (U *UserUsecase) GetUserByID(req *pb.UserIDRequest) (domain.UserDetails, error) {
	userDetails, err := U.UserRepo.GetUserByID(uint(req.UserID))

	if err != nil {
		return userDetails, err
	}

	

	return userDetails, nil

}


func (U *UserUsecase) GetUsersByGender(req *pb.UserGenderRequest) ([]domain.UserDetails, error) {
	userDetails, err := U.UserRepo.GetUsersByGender(req.Gender)

	if err != nil {
		return userDetails, err
	}

	return userDetails, nil

}
