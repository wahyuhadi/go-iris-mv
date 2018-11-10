package main

import (
	"fmt"
	"go-iris-mv/router"
	"os"

	"../go-iris-mv/config"
	"../go-iris-mv/model"
	"github.com/joho/godotenv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func DBMigrate() { /* Auto Migrations */
	fmt.Println("[::] Migration Databases .....")
	db := config.GetDatabaseConnection() /* Get connction to database */
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Profile{}) /* Migration Models */
	fmt.Println("[::] Migration Databases Done")
}

func InitApps() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[-] Error loading .env file")
	}
	fmt.Println("[::] APP Running on Port " + os.Getenv("API_PORT"))
}

func main() {
	app := iris.New()
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
