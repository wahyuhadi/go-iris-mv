package controller

import "github.com/kataras/iris"

func WelcomeController(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  ctx.GetStatusCode(),
		"message": "apps was running",
	})
}
