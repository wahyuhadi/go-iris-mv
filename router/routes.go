//---------------------------------------------------
// function untuk handle router
//---------------------------------------------------

package router

import (
	"go-iris-mv/controller/RequestController"
	"go-iris-mv/controller/UserController"
	"os"

	"../controller"
	"../middleware"

	"github.com/kataras/iris"
)

//---------------------------------------------------
// myXML digunkan untuk generate apidocs
//---------------------------------------------------
// type myXML struct {
// 	Result string `xml:"result"`
// }

func Routers() {
	app := iris.New()
	//---------------------------------------------------
	//  anda bisa menghapus ini jika sudah tahap production
	//	modul ini digunakan untuk generate apidoc
	//---------------------------------------------------
	// yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
	// 	On:       true,
	// 	DocTitle: "Iris",
	// 	DocPath:  "apidocs/index.html",
	// 	BaseUrls: map[string]string{"Production": "", "Staging": "", "Development": "localhost:3000"},
	// })
	// app.Use(irisyaag.New()) // <- IMPORTANT, reg
	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	//---------------------------------------------------
	// for root endpoint (/)
	//---------------------------------------------------
	app.Get("/",
		middleware.WelcomeMiddleware,
		middleware.SecondMiddleware,
		controller.WelcomeController,
	)

	//---------------------------------------------------
	// for v1 endpoint (/v1/user)
	//---------------------------------------------------
	v1 := app.Party("/v1")
	{
		v1.Post("/user", UserController.CreateUser)
		v1.Post("/user/login", UserController.Login)
		v1.Get("/user", UserController.GetAll)
		v1.Get("/get_all", UserController.GetAllUser)
		v1.Get("/user/{id : int}", UserController.GetById)
		v1.Put("/user/{id : int}", UserController.UpdateUser)
		v1.Delete("/user/{id : int}", UserController.DeleteUser)
	}

	//---------------------------------------------------
	// for v1 endpoint (/v1/profile)
	//---------------------------------------------------
	profile := app.Party("/v1/profile")
	{
		profile.Post("/", middleware.DecodeTokenUser, UserController.CreateProfile)
	}

	//---------------------------------------------------
	// for v1 endpoint (/v1/http)
	//---------------------------------------------------
	http := app.Party("/v1/http")
	{
		http.Get("/", RequestController.GetHttpReq)

	}

	app.Run(iris.Addr(":" + os.Getenv("API_PORT"))) // starter handler untuk route
}
