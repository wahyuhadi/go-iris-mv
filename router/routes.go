package router

import (
	"../controller"
	"../middleware"
	"os"

	"github.com/kataras/iris"
)

func Routers() {
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
		v1.Post("/user", controller.CreateUser)
		v1.Post("/user/login", controller.Login)
		v1.Get("/user", controller.GetAll)
		v1.Get("/get_all", controller.GetAllUser)
		v1.Get("/user/{id : int}", controller.GetById)
		v1.Put("/user/{id : int}", controller.UpdateUser)
		v1.Delete("/user/{id : int}", controller.DeleteUser)
	}

	profile := app.Party("/v1/profile")
	{
		profile.Post("/", middleware.DecodeTokenUser, controller.CreateProfile)

	}

	http := app.Party("/v1/http")
	{
		http.Get("/", controller.GetHttpReq)

	}

	app.Run(iris.Addr(":" + os.Getenv("API_PORT"))) // starter handler untuk route
}
