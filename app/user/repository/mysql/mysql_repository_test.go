package mysql

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/choi-yh/example-golang/app/user/model"
	"github.com/choi-yh/example-golang/internal/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	sqlDB *sql.DB
	mock  sqlmock.Sqlmock
	err   error

	db *gorm.DB
	r  *repository
)

func TestMain(m *testing.M) {
	sqlDB, mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db",
		DriverName:                "mysql",
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})

	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r = &repository{
		db: db,
	}

	m.Run()
}

func TestRepository_CreateUser(t *testing.T) {
	ctx := context.Background()

	param := model.CreateUserDBParam{
		ID:        util.CreateID(),
		Email:     "test@test.com",
		Name:      "test_name",
		Phone:     "",
		CreatedAt: util.GetNowPtr(),
	}

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `user` (`email`,`name`,`phone`,`created_at`,`updated_at`,`id`) VALUES (?,?,?,?,?,?)")).
		WithArgs(param.Email, param.Name, param.Phone, param.CreatedAt, sqlmock.AnyArg(), param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = r.CreateUser(ctx, param)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
