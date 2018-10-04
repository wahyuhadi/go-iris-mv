package config

import (
	"log"
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
	// check connectin database
	conn, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONNECTION")) // Connection to database
	if err != nil {
		log.Fatal("Could not connect to the database")
	}

	gormConn = conn
	return gormConn
}
