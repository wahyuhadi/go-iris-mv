package controller

import "github.com/kataras/iris"

func WelcomeController(ctx iris.Context) {
	value := ctx.Values().GetString("status") // call from middleware
	ctx.JSON(iris.Map{
		"status":  ctx.GetStatusCode(),
		"message": "apps was running",
		"info" : value,
	})
}
