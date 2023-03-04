package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConnectDB(t *testing.T) {
	db := ConnectMySQL()

	sqlDB, err := db.DB()
	assert.NoError(t, err)

	err = sqlDB.Ping()
	assert.NoError(t, err)
}
