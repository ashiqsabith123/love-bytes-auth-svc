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
