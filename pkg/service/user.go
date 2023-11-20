package service

import (
	"context"

	"github.com/ashiqsabith123/auth-svc/pkg/domain"
	usecase "github.com/ashiqsabith123/auth-svc/pkg/usecase/interfaces"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
	"github.com/jinzhu/copier"
)

type UserService struct {
	UserUsecase usecase.UserUsecase
	pb.UnimplementedAuthServiceServer
}

func NewUserService(usecase usecase.UserUsecase) UserService {
	return UserService{UserUsecase: usecase}
}

func (U *UserService) Signup(ctx context.Context, req *pb.SignUpReq) (*pb.Responce, error) {
	var user domain.User

	copier.Copy(user, &req)

	U.UserUsecase.SignUp(user)

	return &pb.Responce{
		Message: "Login succesfull",
		Code:    200,
	}, nil
}

func (U *UserService) SendOtp(ctx context.Context, req *pb.OtpReq) (*pb.Responce, error) {

	resp, err := U.UserUsecase.SendOtp(req.Phone)

	if err != nil {
		return &pb.Responce{
			Code:    500,
			Message: resp,
			Error:   err.Error(),
		}, nil
	}

	return &pb.Responce{
		Code:    200,
		Message: resp,
	}, nil
}
