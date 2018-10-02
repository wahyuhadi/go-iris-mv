package controller

import (
	"github.com/kataras/iris"
	"go-iris-mv/model"
)

func (idb* InDB) CreteUser(ctx iris.Context) {
	var (
		user model.User
	)
	ctx.ReadJSON(&user)
	idb.DB.Create(&user)
	ctx.JSON(iris.Map{
		"error" : "false",
		"status" : iris.StatusOK,
		"result" : user,
	})
}

func (idb* InDB) GetAll(ctx iris.Context)  {
	var(
		user []model.User // [] for array result
		result  iris.Map
	)

	ctx.ReadJSON(&user)
	idb.DB.Find(&user)
	if len(user) <= 0 {
		result = iris.Map{
			"error" : "false",
			"status" : iris.StatusOK,
			"result" : nil,
			"count" : 0,
		}
	} else {
		result = iris.Map{
			"error" : "false",
			"status" : iris.StatusOK,
			"result" : user,
			"count" : len(user),
		}
	}
	ctx.JSON(result)
}

func  (idb* InDB) GetById(ctx iris.Context)  {
	var (
		user model.User
		result iris.Map
	)

	id := ctx.Params().Get("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if  err != nil {
		result = iris.Map{
			"error" : "true",
			"status" : iris.StatusBadRequest,
			"result" : err.Error(),
			"count" : 0,
		}
	} else {
		result = iris.Map{
			"error" : "false",
			"status" : iris.StatusOK,
			"result" : user,
			"count" : 1,
		}
	}
	ctx.JSON(result)
}

func (idb* InDB) UpdateUser (ctx iris.Context)  {
	var (
		user model.User
		newUser model.User
		result iris.Map
	)
	id := ctx.Params().Get("id")
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = iris.Map{
			"error" : "true",
			"status" : iris.StatusBadRequest,
			"message" : "user not found",
			"result" : nil,
		}
	}
	ctx.ReadJSON(&newUser)
	err = idb.DB.Model(&user).Updates(newUser).Error
	if err != nil {
		result = iris.Map{
			"error" : "true",
			"status" : iris.StatusBadRequest,
			"message" : "error when update user",
			"result" : err.Error(),
		}
	} else {
		result = iris.Map{
			"error" : "false",
			"status" : iris.StatusOK,
			"message" : "success update user",
			"result" : newUser,
		}
	}
	ctx.JSON(result)
}

func (idb* InDB)  DeleteUser(ctx iris.Context) {
	var(
		user model.User
		result iris.Map
	)
	id := ctx.Params().Get("id")
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = iris.Map{
			"error" : "true",
			"status" : iris.StatusBadRequest,
			"message" : "User not found",
			"result" : nil,
		}
	}

	err = idb.DB.Where("id = ?", id).Delete(&user, id).Error
	if err != nil {
		result = iris.Map{
			"error" : "true",
			"status" : iris.StatusBadRequest,
			"message" : "Failed Delete user",
			"result" : err.Error(),
		}
	} else {
		result = iris.Map{
			"error" : "false",
			"status" : iris.StatusOK,
			"message" : "Failed Delete user",
			"result" : nil,
		}
	}
	ctx.JSON(result)
}