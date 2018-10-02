package middleware

import "github.com/kataras/iris"

func  WelcomeMiddleware(ctx iris.Context)  {
	shareStatus := "welcome to iris golang starter kit"
	ctx.Values().Set("status", shareStatus) // value  yang akan di gunakan dimain handler
	ctx.Next()
}