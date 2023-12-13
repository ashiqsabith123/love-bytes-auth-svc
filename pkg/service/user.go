package service

import (
	"context"
	"fmt"

	usecase "github.com/ashiqsabith123/auth-svc/pkg/usecase/interfaces"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type UserService struct {
	UserUsecase usecase.UserUsecase
	pb.UnimplementedAuthServiceServer
}

func NewUserService(usecase usecase.UserUsecase) UserService {
	return UserService{UserUsecase: usecase}
}

func (U *UserService) SendOtp(ctx context.Context, req *pb.OtpReq) (*pb.Response, error) {

	err := U.UserUsecase.SendOtp(req.Phone)

	if err != nil {
		return &pb.Response{
			Code:    403,
			Message: "Api error",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}
	return &pb.Response{
		Code:    200,
		Message: "Otp send succesfully",
	}, nil
}

func (U *UserService) VerifyOtpAndAuth(ctx context.Context, req *pb.VerifyOtpReq) (*pb.Response, error) {

	token, userFound, status, err := U.UserUsecase.VerifyOtpAndAuth(req)

	if err != nil {
		return &pb.Response{
			Code:    int32(status),
			Message: "Authentication failed",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}

	data := &pb.TokenResp{
		Userfound: userFound,
		Token:     token,
	}

	dataInBytes, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)

	}

	return &pb.Response{
		Code:    int32(status),
		Message: "Auth succesfull",
		Data: &anypb.Any{
			Value: dataInBytes,
		},
	}, nil

}

func (U *UserService) SaveUserDetais(ctx context.Context, req *pb.UserDetailsReq) (*pb.Response, error) {

	err := U.UserUsecase.SaveUserDetails(req)

	if err != nil {
		return &pb.Response{
			Code:    500,
			Message: "Failed to add details",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}

	return &pb.Response{
		Code:    201,
		Message: "Details added succesfully",
	}, nil
}
