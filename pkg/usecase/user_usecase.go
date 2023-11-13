package usecase

import (
	"log"

	"github.com/ashiqsabith123/auth-svc/pkg/domain"
	repo "github.com/ashiqsabith123/auth-svc/pkg/repository/interface"
	usecase "github.com/ashiqsabith123/auth-svc/pkg/usecase/interfaces"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) usecase.UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (U *UserUsecase) SignUp(user domain.User) {
	ok, _ := U.UserRepo.FindUser(user.Phone)

	if !ok {
		err := U.UserRepo.CreateUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}

}
