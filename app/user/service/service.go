package service

import (
	"context"

	"github.com/choi-yh/example-golang/app/user/model"
	"github.com/choi-yh/example-golang/app/user/repository/mysql"
	domainModel "github.com/choi-yh/example-golang/domain/model"
	"github.com/choi-yh/example-golang/internal/util"
)

type service struct {
	mysqlRepository mysql.Repository
}

type Service interface {
	SignUp(ctx context.Context, param model.SignUpParam) (domainModel.User, error)
}

func NewService() Service {
	mysqlRepository := mysql.NewRepository()

	return &service{
		mysqlRepository: mysqlRepository,
	}
}

func (s service) SignUp(ctx context.Context, param model.SignUpParam) (res domainModel.User, err error) {
	// create user
	user := model.CreateUserDBParam{
		ID:        util.CreateID(),
		Email:     param.Email,
		Name:      param.Name,
		Phone:     param.Phone,
		CreatedAt: util.GetNowPtr(),
	}

	err = s.mysqlRepository.CreateUser(ctx, user)
	if err != nil {
		return res, err
	}

	//create user_password

	return
}
