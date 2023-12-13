package interfaces

import "github.com/ashiqsabith123/love-bytes-proto/auth/pb"

type UserUsecase interface {
	SendOtp(phone string) error
	VerifyOtpAndAuth(req *pb.VerifyOtpReq) (string, bool, int, error)
	SaveUserDetails(req *pb.UserDetailsReq) error
}
