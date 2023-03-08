package mysql

import (
	"context"

	"github.com/choi-yh/example-golang/app/user/model"
	domainModel "github.com/choi-yh/example-golang/domain/model"
	utilMysql "github.com/choi-yh/example-golang/internal/util/database/mysql"
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
	data := domainModel.User{
		ID:        param.ID,
		Email:     param.Email,
		Name:      param.Name,
		Phone:     param.Phone,
		CreatedAt: param.CreatedAt,
	}

	if err := r.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
