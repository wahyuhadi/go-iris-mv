package controller

import (
	"encoding/json"
	"go-iris-mv/config"
	"go-iris-mv/model"
	"go-iris-mv/service"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/structs"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
)

// Create user
func CreateUser(ctx iris.Context) {
	var (
		user model.User
	)
	ctx.ReadJSON(&user)
	hash, _ := service.HashPassword(user.Password)
	user.Password = hash
	user.Role = "user"
	db := config.GetDatabaseConnection()
	defer db.Close() // close connecion database to save memory
	db.Create(&user)
	ctx.JSON(iris.Map{
		"error":  "false",
		"status": iris.StatusOK,
		"result": user,
	})
}

// Login
func Login(ctx iris.Context) {
	var (
		user   model.User
		result iris.Map
	)
	ctx.ReadJSON(&user)
	email := user.Email
	pass := user.Password
	db := config.GetDatabaseConnection()
	defer db.Close()
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "Invalid login credentials. Please try again",
		}
		ctx.JSON(result)
		return
	}

	// if email not found
	if user.ID == 0 {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "Invalid login credentials. Please try again",
		}
		ctx.JSON(result)
		return
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

	// compare password
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

	// To generate JWT token
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
			"role":    user.Role,
		}
	}
	ctx.JSON(result)
	return
}

// Get all user
func GetAll(ctx iris.Context) {
	var (
		user   []model.User // [] for array result
		result iris.Map
	)

	ctx.ReadJSON(&user)
	db := config.GetDatabaseConnection()
	defer db.Close()
	db.Preload("Profile").Find(&user) // relation to profile, You can fix this ??
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

// Get user by id
func GetById(ctx iris.Context) {
	var (
		user   model.User
		result iris.Map
		data   map[string]interface{}
	)

	id := ctx.Params().Get("id")
	db := config.GetDatabaseConnection()
	defer db.Close()
	err := db.Where("id = ?", id).Preload("Profile").First(&user).Error
	mapstruct := structs.Map(&user)
	delete(mapstruct, "Profile")
	delete(mapstruct, "Password")
	mar, _ := json.Marshal(mapstruct)
	byt := []byte(strings.ToLower(string(mar)))
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

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
			"data":   data,
			"count":  1,
		}
	}
	ctx.JSON(result)
	return
}

// Update user by id
func UpdateUser(ctx iris.Context) {
	var (
		user    model.User
		newUser model.User
		result  iris.Map
	)
	id := ctx.Params().Get("id") // get id by params
	db := config.GetDatabaseConnection()
	defer db.Close()
	err := db.First(&user, id).Error
	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "user not found",
			"result":  nil,
		}
	}
	ctx.ReadJSON(&newUser)
	err = db.Model(&user).Updates(newUser).Error
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
	return
}

// Delete user by id
func DeleteUser(ctx iris.Context) {
	var (
		user   model.User
		result iris.Map
	)
	id := ctx.Params().Get("id") // get id by params
	db := config.GetDatabaseConnection()
	defer db.Close()
	err := db.First(&user, id).Error
	if err != nil {
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "User not found",
			"result":  nil,
		}
	}

	err = db.Where("id = ?", id).Delete(&user, id).Error
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
	return
}

// create profile
func CreateProfile(ctx iris.Context) {
	var (
		profile model.Profile
	)

	id := ctx.Values().Get("id") // get id from middleware
	ctx.ReadJSON(&profile)
	var userId int64
	userId = int64(id.(float64)) // convertion type float64 to int64
	profile.UserID = userId
	db := config.GetDatabaseConnection()
	defer db.Close()
	db.Create(&profile)
	ctx.JSON(iris.Map{
		"error":  "false",
		"status": iris.StatusOK,
		"result": profile,
	})
	return
}
