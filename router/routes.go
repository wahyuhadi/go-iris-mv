package router

import (
	"github.com/kataras/iris"
	"go-iris-mv/config"
	"go-iris-mv/controller"
	"go-iris-mv/middleware"
	"os"
)

func Routers() {
	db := config.GetDatabaseConnection()
	inDB := &controller.InDB{DB: db}
	app := iris.Default()
	// for / endpoint
	app.Get("/",
		middleware.WelcomeMiddleware,
		middleware.SecondMiddleware,
		controller.WelcomeController,
	)

	// example group: v1
	v1 := app.Party("/v1")
	{
		v1.Post("/user", inDB.CreteUser)
		v1.Post("/user/login", inDB.Login)
		v1.Get("/user", inDB.GetAll)
		v1.Get("/user/{id : int}", inDB.GetById)
		v1.Put("/user/{id : int}", inDB.UpdateUser)
		v1.Delete("/user/{id : int}", inDB.DeleteUser)
	}

	profile := app.Party("/v1/profile")
	{
		profile.Post("/", middleware.DecodeTokenUser, inDB.CreateProfile)

	}

	app.Run(iris.Addr(":" + os.Getenv("API_PORT"))) // starter handler untuk route
}
