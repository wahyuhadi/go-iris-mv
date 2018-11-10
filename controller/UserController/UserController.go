package UserController

import (
	"os"
	"time"

	"go-iris-mv/config"
	UserQuery "go-iris-mv/controller/UserController/UseCase"
	"go-iris-mv/model"
	"go-iris-mv/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
)

/*********************************************************
| Created By : @wahyuhadi
| Desc : Create Users
**********************************************************/
func CreateUser(ctx iris.Context) {
	var (
		user model.User
	)
	ctx.ReadJSON(&user)
	hash, _ := service.HashPassword(user.Password) /*  generate hash with bycryp */
	user.Password = hash
	user.Role = "user"                   /*  add user role */
	db := config.GetDatabaseConnection() /*  Open connectins */
	defer db.Close()
	db.Create(&user)
	ctx.JSON(iris.Map{
		"error":  "false",
		"status": iris.StatusOK,
		"result": user,
	})
	return
}

/*********************************************************
| Created By : @wahyuhadi
| Desc : Login
**********************************************************/
func Login(ctx iris.Context) {
	var (
		user   model.User
		result iris.Map
	)
	ctx.ReadJSON(&user)
	email := user.Email /* get Email from body */
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

	/*  Email Not Found */
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

	/*  Compare password */
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { /*  Password tidak sesuai */
		result = iris.Map{
			"error":   "true",
			"status":  iris.StatusBadRequest,
			"message": "Invalid login credentials. Please try again",
		}
		ctx.JSON(result)
		return
	}

	/*  Generate JWT token */
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
			"role":    user.Role, /*  token  */
		}
	}
	ctx.JSON(result)
	return
}

/*********************************************************
| Created By : @wahyuhadi
| Desc : get all user menggunakan self qeury
**********************************************************/
func GetAllUser(ctx iris.Context) {

	/*********************************************************
	| Created By : @wahyuhadi
	| Desc : Profile and user struct
	**********************************************************/
	/*  Profile struct */
	type Profile struct {
		ID        int64      `json:"id" gorm:"primary_key"`
		UserID    int64      `json:"user_id,omitempty" gorm:"type:bigint REFERENCES users(id)"`
		Address   string     `json:"address,omitempty" gorm:"not null; type:varchar(100)"`
		LastName  string     `json:"lastname,omitempty" gorm:"not null; type:varchar(100)"`
		FirstName string     `json:"firstname,omitempty" gorm:"not null; type:varchar(100)"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
	}

	/*  User */
	type User struct {
		ID        int64      `json:"id" gorm:"primary_key"`
		Role      string     `json:"role,omitempty" gorm:"not null; type:ENUM('admin', 'user', 'root')"`
		Email     string     `json:"email" gorm:"not null; size:255"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
		Profile   Profile    `json:"profile"` /*  get from model */
	}

	var (
		users  []User
		result iris.Map
	)

	err := UserQuery.GetAllAssociation(&users, "Profile") // pemaggilan service self query untuk association dengan profile
	if err != nil {
		result = iris.Map{
			"error":  "true",
			"status": iris.StatusBadRequest,
			"result": nil,
		}
	} else {
		result = iris.Map{
			"error":  "false",
			"status": iris.StatusOK,
			"result": users,
			"count":  len(users),
		}
	}
	ctx.JSON(result)
	return
}

/*********************************************************
| Created By : @wahyuhadi
| Desc : Get all User witj Global Query
**********************************************************/
func GetAll(ctx iris.Context) {

	type User struct {
		ID        int64      `json:"id" gorm:"primary_key"`
		Role      string     `json:"role,omitempty" gorm:"not null; type:ENUM('admin', 'user', 'root')"`
		Email     string     `json:"email" gorm:"not null; size:255"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
		//Profile   Profile    `json:"profile"` // Get from model->Profile
	}

	var (
		users  []User /*  Array result */
		result iris.Map
	)

	err := service.GetAll(&users)
	if err != nil {
		result = iris.Map{
			"error":  "true",
			"status": iris.StatusBadRequest,
			"result": nil,
		}
	} else {
		result = iris.Map{
			"error":  "false",
			"status": iris.StatusOK,
			"result": users,
			"count":  len(users),
		}
	}

	ctx.JSON(result)
	return
}

/*********************************************************
| Created By : @wahyuhadi
| Desc : Get user by Id
**********************************************************/
func GetById(ctx iris.Context) {
	var (
		users  model.User
		result iris.Map
	)

	id := ctx.Params().Get("id")
	err := service.GetById(&users, id)
	if err != nil {
		result = iris.Map{
			"error":  "true",
			"status": iris.StatusBadRequest,
			"result": nil,
		}
	} else {
		result = iris.Map{
			"error":  "false",
			"status": iris.StatusOK,
			"result": users,
			"count":  1,
		}
	}
	ctx.JSON(result)
	return
}

/*********************************************************
| Created By : @wahyuhadi
| Desc : Update User
**********************************************************/
func UpdateUser(ctx iris.Context) {
	var (
		user    model.User
		newUser model.User
		result  iris.Map
	)
	id := ctx.Params().Get("id") /*  get id in params */
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

/*********************************************************
| Created By : @wahyuhadi
| Desc : Delete user By id
**********************************************************/
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

/*********************************************************
| Created By : @wahyuhadi
| Desc : Create Profile
**********************************************************/
func CreateProfile(ctx iris.Context) {
	type Profile struct {
		ID        int64      `json:"id" gorm:"primary_key"`
		UserID    int64      `json:"user_id,omitempty" gorm:"type:bigint REFERENCES users(id)"`
		Address   string     `json:"address,omitempty" gorm:"not null; type:varchar(100)"`
		LastName  string     `json:"lastname,omitempty" gorm:"not null; type:varchar(100)"`
		FirstName string     `json:"firstname,omitempty" gorm:"not null; type:varchar(100)"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
		//User      User       `gorm:"foreignkey:UserRefer"`
	}
	var (
		profile Profile
	)

	id := ctx.Values().Get("id") // get id from middleware
	ctx.ReadJSON(&profile)
	var userID int64
	userID = int64(id.(float64)) // contoh perubahan type dari integer ke int64
	profile.UserID = userID
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
