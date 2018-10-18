package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gormConn *gorm.DB

// get connection database
func GetDatabaseConnection() *gorm.DB { // Check Connection Status
	if gormConn != nil && gormConn.DB() != nil && gormConn.DB().Ping() == nil {
		return gormConn
	}

	conn, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONNECTION")) // Connection to database
	conn.DB().SetMaxOpenConns(200)                                              // Sane default
	conn.DB().SetMaxIdleConns(10)
	//conn.DB().SetConnMaxLifetime(time.Nanosecond)
	if err != nil {
		panic("Could not connect to the database") // log error without close
	}
	conn.LogMode(true)

	gormConn = conn
	return gormConn
}
