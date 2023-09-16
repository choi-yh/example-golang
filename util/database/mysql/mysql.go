package mysql

import (
	"fmt"
	"log"
	"os"

	ppgorm "github.com/pinpoint-apm/pinpoint-go-agent/plugin/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	user     = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	host     = os.Getenv("MYSQL_HOST")
	port     = os.Getenv("MYSQL_PORT")
	database = os.Getenv("MYSQL_DATABASE")
)

func ConnectMySQL() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	config := &gorm.Config{}

	db, err := ppgorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Panic(err)
	}

	return db
}
