package main

import (
	"fmt"
	"os"

	"../go-iris-mv/config"
	"../go-iris-mv/model"
	"../go-iris-mv/router"
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/joho/godotenv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

type myXML struct {
	Result string `xml:"result"`
}

func DBMigrate() { // auto migration
	fmt.Println("[::] Migration Databases .....")
	db := config.GetDatabaseConnection() // check connection to Databases
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Profile{}) // Migrate Model
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
	app := iris.Default()
	InitApps()
	DBMigrate()

	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Iris",
		DocPath:  "apidocs/apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New()) // <- IMPORTANT, reg
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
