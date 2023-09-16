package service

import (
	"context"
	"fmt"

	domainModel "github.com/choi-yh/example-golang/application/domain"
	model2 "github.com/choi-yh/example-golang/application/user/model"
	"github.com/choi-yh/example-golang/application/user/repository/mysql"
	"github.com/choi-yh/example-golang/util"
	"github.com/google/uuid"
	"github.com/pinpoint-apm/pinpoint-go-agent"
)

type service struct {
	mysqlRepository mysql.Repository
}

type Service interface {
	SignUp(ctx context.Context, param model2.SignUpParam) (domainModel.User, error)
}

func NewService() Service {
	mysqlRepository := mysql.NewRepository()

	return &service{
		mysqlRepository: mysqlRepository,
	}
}

func (s service) SignUp(ctx context.Context, param model2.SignUpParam) (res domainModel.User, err error) {
	defer pinpoint.FromContext(ctx).NewSpanEvent("SignUp").EndSpanEvent()

	hashPassword, err := util.EncodeHash(param.Password)
	if err != nil {
		return res, err
	}

	if param.Name == "error" {
		return res, fmt.Errorf("name error")
	}

	user := model2.CreateUserDBParam{
		ID:        uuid.NewString(),
		Email:     param.Email,
		Password:  hashPassword,
		Name:      param.Name,
		Phone:     param.Phone,
		CreatedAt: util.GetNow(),
	}

	err = s.mysqlRepository.CreateUser(ctx, user)
	if err != nil {
		return res, err
	}

	res = domainModel.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
	}

	return
}
