package controller

import "github.com/kataras/iris"

// Welcome Controller
func WelcomeController(ctx iris.Context) {
	value := ctx.Values().GetString("status") // call from middleware
	author := ctx.Values().GetString("author") // call from middleware
	ctx.JSON(iris.Map{
		"status":  ctx.GetStatusCode(),
		"message": "apps running",
		"info" : value,
		"author" : author,
	})
}
