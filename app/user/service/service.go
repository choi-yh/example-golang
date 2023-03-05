package service

import "github.com/choi-yh/example-golang/app/user/repository/mysql"

type service struct {
	mysqlRepository mysql.Repository
}

type Service interface {
}

func NewService() Service {
	mysqlRepository := mysql.NewRepository()

	return &service{
		mysqlRepository: mysqlRepository,
	}
}
