package middleware

import "github.com/kataras/iris"

// First Middleware
func WelcomeMiddleware(ctx iris.Context) {
	shareStatus := "welcome to iris golang starter kit"
	ctx.Values().Set("status", shareStatus) // value  yang akan di gunakan dimain handler
	ctx.Next()                              // execute next handle , pada kasus ini adalah Welcome Controller
}

// second middleware
func SecondMiddleware(ctx iris.Context) {
	SecondMiddleware := "Create by Rahmat Wahyu Hadi"
	ctx.Values().Set("author", SecondMiddleware)
	ctx.Next()
}
