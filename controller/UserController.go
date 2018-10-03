package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"go-iris-mv/model"
	"go-iris-mv/service"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func (idb *InDB) CreteUser(ctx iris.Context) {
	var (
		user model.User
	)
	ctx.ReadJSON(&user)
	hash, _ := service.HashPassword(user.Password)
	user.Password = hash
	user.Role = "user"
	idb.DB.Create(&user)
	ctx.JSON(iris.Map{
		"error":  "false",
		"status": iris.StatusOK,
		"result": user,
	})
}

func (idb *InDB) Login(ctx iris.Context) {
	var (
		user   model.User
		result iris.Map
	)
	ctx.ReadJSON(&user)
	email := user.Email
	pass := user.Password
	err := idb.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "Invalid login credentials. Please try again",
		}
	}

	if user.Role == "admin" {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "Invalid login credentials. Please try again",
		}
		ctx.JSON(result)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "Invalid login credentials. Please try again",
		}
		ctx.JSON(result)
		return
	}

	user.Password = ""
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["role"] = user.Role
	sign.Claims = claims
	token, err := sign.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": err.Error(),
		}
	} else {
		result = iris.Map{
			"error":   "false",
			"status":  iris.StatusOK,
			"message": "success login",
			"token":   token,
		}
	}
	ctx.JSON(result)
	return
}

func (idb *InDB) GetAll(ctx iris.Context) {
	var (
		user    []model.User // [] for array result
		profile model.Profile
		result  iris.Map
	)

	ctx.ReadJSON(&user)
	idb.DB.Find(&user).Related(&profile)
	if len(user) <= 0 {
		result = iris.Map{
			"error":  "false",
			"status": iris.StatusOK,
			"result": nil,
			"count":  0,
		}
	} else {
		result = iris.Map{
			"error":  "false",
			"status": iris.StatusOK,
			"result": user,
			"count":  len(user),
		}
	}
	ctx.JSON(result)
	return
}

func (idb *InDB) GetById(ctx iris.Context) {
	var (
		user   model.User
		result iris.Map
	)

	id := ctx.Params().Get("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		result = iris.Map{
			"error":  "true",
			"status": iris.StatusBadRequest,
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = iris.Map{
			"error":  "false",
			"status": iris.StatusOK,
			"result": user,
			"count":  1,
		}
	}
	ctx.JSON(result)
}

func (idb *InDB) UpdateUser(ctx iris.Context) {
	var (
		user    model.User
		newUser model.User
		result  iris.Map
	)
	id := ctx.Params().Get("id")
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "user not found",
			"result":  nil,
		}
	}
	ctx.ReadJSON(&newUser)
	err = idb.DB.Model(&user).Updates(newUser).Error
	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "error when update user",
			"result":  err.Error(),
		}
	} else {
		result = iris.Map{
			"error":   "false",
			"status":  iris.StatusOK,
			"message": "success update user",
			"result":  newUser,
		}
	}
	ctx.JSON(result)
}

func (idb *InDB) DeleteUser(ctx iris.Context) {
	var (
		user   model.User
		result iris.Map
	)
	id := ctx.Params().Get("id")
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "User not found",
			"result":  nil,
		}
	}

	err = idb.DB.Where("id = ?", id).Delete(&user, id).Error
	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "Failed Delete user",
			"result":  err.Error(),
		}
	} else {
		result = iris.Map{
			"error":   "false",
			"status":  iris.StatusOK,
			"message": "Failed Delete user",
			"result":  nil,
		}
	}
	ctx.JSON(result)
}

func (idb *InDB) CreateProfile(ctx iris.Context) {
	var (
		profile model.Profile
	)

	//id := ctx.Values().Get("id") // get id from middleware
	ctx.ReadJSON(&profile)

	profile.UserID = 1 // need fixing how to change float64 to int
	idb.DB.Create(&profile)
	ctx.JSON(iris.Map{
		"error":  "false",
		"status": iris.StatusOK,
		"result": profile,
	})

}
