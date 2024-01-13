package interfaces

import (
	"github.com/ashiqsabith123/auth-svc/pkg/domain"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
)

type UserUsecase interface {
	SendOtp(phone string) error
	VerifyOtpAndAuth(req *pb.VerifyOtpReq) (string, bool, int, error)
	SaveUserDetails(req *pb.UserDetailsReq) error
	GetUserByID(req *pb.UserIDRequest) (domain.UserDetails, error)
	GetUsersByGender(req *pb.UserGenderRequest) ([]domain.UserDetails, error)
}
