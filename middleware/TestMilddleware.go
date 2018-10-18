//---------------------------------------------------
// Contoh  penerapan middleware
//---------------------------------------------------

package middleware

import "github.com/kataras/iris"

//---------------------------------------------------
// first middleware, bisa dilihat pada welcome controler untuk get value
//---------------------------------------------------
func WelcomeMiddleware(ctx iris.Context) {
	shareStatus := "welcome to iris golang starter kit"
	ctx.Values().Set("status", shareStatus) // value  yang akan di gunakan dimain handler
	ctx.Next()                              // execute next handle , pada kasus ini adalah Welcome Controller
}

//---------------------------------------------------
// first middleware, bisa dilihat pada welcome controler untuk get value
//---------------------------------------------------
func SecondMiddleware(ctx iris.Context) {
	SecondMiddleware := "Create by Rahmat Wahyu Hadi"
	ctx.Values().Set("author", SecondMiddleware)
	ctx.Next()
}
