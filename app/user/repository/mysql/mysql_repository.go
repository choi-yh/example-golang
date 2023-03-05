package mysql

import (
	utilMysql "github.com/choi-yh/example-golang/internal/util/database/mysql"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
}

func NewRepository() Repository {
	db := utilMysql.ConnectMySQL()

	return &repository{
		db: db,
	}
}
