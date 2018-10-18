package router

import (
	"os"

	"../controller"
	"../middleware"

	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
)

// ini digunakan untuk saat development
type myXML struct {
	Result string `xml:"result"`
}

func Routers() {
	app := iris.New()
	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	/***
		anda bisa menghapus ini jika sudah tahap production
		modul ini digunakan untuk generate apidoc
	***/
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Iris",
		DocPath:  "apidocs/index.html",
		BaseUrls: map[string]string{"Production": "", "Staging": "", "Development": "localhost:3000"},
	})
	app.Use(irisyaag.New()) // <- IMPORTANT, reg
	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

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
