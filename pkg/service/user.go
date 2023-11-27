package service

import (
	"context"

	usecase "github.com/ashiqsabith123/auth-svc/pkg/usecase/interfaces"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
)

type UserService struct {
	UserUsecase usecase.UserUsecase
	pb.UnimplementedAuthServiceServer
}

func NewUserService(usecase usecase.UserUsecase) UserService {
	return UserService{UserUsecase: usecase}
}

func (U *UserService) Signup(ctx context.Context, req *pb.OtpSignUpReq) (*pb.Responce, error) {

	resp, status, err := U.UserUsecase.VerifyOtpAndSignUp(req)

	var code int32

	switch status {
	case 1:
		code = 200
	case 2:
		code = 401
	case 3:
		code = 403
	case 5:
		code = 500
	case 6:
		code = 400

	}

	if err != nil {
		return &pb.Responce{
			Message: resp,
			Code:    code,
			Error:   err.Error(),
		}, nil
	}

	return &pb.Responce{
		Message: "Signup succesfull",
		Code:    200,
	}, nil
}

func (U *UserService) SendOtp(ctx context.Context, req *pb.OtpReq) (*pb.Responce, error) {

	resp, status, err := U.UserUsecase.SendOtp(req.Phone)

	var code int32

	switch status {
	case 1:
		code = 500
	case 3:
		code = 500
	case 2:
		code = 400

	}

	if err != nil {
		return &pb.Responce{
			Code:    code,
			Message: resp,
			Error:   err.Error(),
		}, nil
	}
	return &pb.Responce{
		Code:    200,
		Message: resp,
	}, nil
}
