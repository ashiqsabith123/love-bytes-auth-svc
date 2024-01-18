package service

import (
	"context"
	"fmt"
	"net/http"

	usecase "github.com/ashiqsabith123/auth-svc/pkg/usecase/interfaces"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
	logs "github.com/ashiqsabith123/love-bytes-proto/log"
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
			Code:    http.StatusForbidden,
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
		logs.ErrLog.Println(err)
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
		return &pb.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed while marshaling",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
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

func (U *UserService) GetUserByID(ctx context.Context, req *pb.UserIDRequest) (*pb.Response, error) {

	userDetails, err := U.UserUsecase.GetUserByID(req)

	if err != nil {
		return &pb.Response{
			Code:    500,
			Message: "Failed to get user details",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}

	fmt.Println(userDetails)

	data := &pb.UserRepsonse{
		UserID:   int32(userDetails.UserID),
		Fullname: userDetails.Fullname,
		Location: userDetails.Location,
		Dob:      userDetails.DateOfBirth,
		Lat:      userDetails.Latitude,
		Log:      userDetails.Longitude,
		Gender:   userDetails.Gender,
	}

	dataInBytes, err := proto.Marshal(data)
	if err != nil {
		return &pb.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed while marshaling",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}

	return &pb.Response{
		Code:    http.StatusOK,
		Message: "Data fetched succesfully",
		Data: &anypb.Any{
			Value: dataInBytes,
		},
	}, nil
}

func (U *UserService) GetUsersByGender(ctx context.Context, req *pb.UserGenderRequest) (*pb.Response, error) {
	resp, err := U.UserUsecase.GetUsersByGender(req)

	if err != nil {
		return &pb.Response{
			Code:    500,
			Message: "Failed to get users by gender",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}

	data := make([]*pb.UserRepsonse, len(resp))

	for i, v := range resp {
		userData := &pb.UserRepsonse{
			UserID:   int32(v.UserID),
			Fullname: v.Fullname,
			Location: v.Location,
			Lat:      v.Latitude,
			Log:      v.Longitude,
			Dob:      v.DateOfBirth,
		}

		data[i] = userData
	}

	userData := pb.UserResponses{
		UserRepsonses: data,
	}

	dataInBytes, err := proto.Marshal(&userData)
	if err != nil {
		return &pb.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed while marshaling",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}

	return &pb.Response{
		Code:    http.StatusOK,
		Message: "Data fetched succesfully",
		Data: &anypb.Any{
			Value: dataInBytes,
		},
	}, nil
}

