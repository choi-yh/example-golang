package mysql

import (
	"context"

	domainModel "github.com/choi-yh/example-golang/application/domain"
	"github.com/choi-yh/example-golang/application/user/model"
	utilMysql "github.com/choi-yh/example-golang/util/database/mysql"
	"github.com/pinpoint-apm/pinpoint-go-agent"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateUser(ctx context.Context, param model.CreateUserDBParam) error
}

func NewRepository() Repository {
	db := utilMysql.ConnectMySQL()

	return &repository{
		db: db,
	}
}

func (r repository) CreateUser(ctx context.Context, param model.CreateUserDBParam) error {
	defer pinpoint.FromContext(ctx).NewSpanEvent("CreateUser").EndSpanEvent()

	data := domainModel.User{
		ID:        param.ID,
		Email:     param.Email,
		Password:  param.Password,
		Name:      param.Name,
		Phone:     param.Phone,
		CreatedAt: param.CreatedAt,
	}

	r.db = r.db.WithContext(ctx)
	if err := r.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
