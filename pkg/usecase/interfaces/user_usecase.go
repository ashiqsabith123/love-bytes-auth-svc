package interfaces

import (
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
)

type UserUsecase interface {
	VerifyOtpAndSignUp(rep *pb.OtpSignUpReq) (string, int, error)
	SendOtp(phone string) (string, int, error)
}
