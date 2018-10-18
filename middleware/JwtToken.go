package middleware

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

//---------------------------------------------------
// Decode user Token
//---------------------------------------------------
func DecodeTokenUser(ctx iris.Context) {
	var (
		result iris.Map
	)

	token := ctx.GetHeader("token")
	if token == "" {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "token not found",
		}
		ctx.JSON(result)
		return
	}

	decode, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil // get from .env
	})

	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "token invalid",
		}
		ctx.JSON(result)
		return
	}

	claims := decode.Claims.(jwt.MapClaims)
	for key, val := range claims {
		ctx.Values().SetImmutable(key, val) //value  number has been change with float64
	}
	ctx.Next() //Store  value ke next handler
}
