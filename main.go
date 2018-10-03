package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-iris-mv/config"
	"go-iris-mv/model"
	"go-iris-mv/router"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func DBMigrate() { // auto migration
	fmt.Println("[::] Migration Databases .....")
	db := config.GetDatabaseConnection() // check connection to Databases
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Profile{})  // Migrate Model
	fmt.Println("[::] Migration Databases Done")
}

func InitApps() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("[::] APP Running on Port " + os.Getenv("API_PORT"))
}

func main() {
	app := iris.Default()
	InitApps()
	DBMigrate()
	requestLogger := logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		MessageContextKeys: []string{"logger_message"},
		MessageHeaderKeys:  []string{"User-Agent"},
	})
	app.Use(requestLogger)
	router.Routers()
}

